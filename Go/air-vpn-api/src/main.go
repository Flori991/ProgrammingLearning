package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Flori991/ProgrammingLearning/cache"
)

func main() {
	logStartup("Starting up airvpn-api.")
	initialize()
	startServer()
}

func initialize() {
	// Initialize configuration
	initConfig()
	// Initialize cache
	initCache()
}

func initCache() {
	logStartup("Initializing Cache.")
	appCache = cache.NewCache(config.CacheTtl)
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/dashboard", handleDashboardData)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: mux,
	}

	// Configure graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		logStartup("Server starting on port: " + config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logError("Server failed to start", err)
			os.Exit(1)
		}
	}()

	logStartup("Server ready")

	// Wait for interrupt signal to gracefully shutdown the server
	<-stop

	logStartup("Shutdown signal received, gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logStartup("Server forced to shutdown")
	} else {
		logStartup("Server stopped cleanly")
	}
}
