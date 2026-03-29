package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Flori991/ProgrammingLearning/cache"
)

var appCache *cache.Cache

func main() {
	initialize()
	startServer()
}

func initialize() {
	logStartup("Starting up airvpn-api.")

	// Initialize logger
	logStartup("Initializing Log Level.")
	initLogLevel()

	// Initialize cache
	logStartup("Initializing Cache.")
	initCache()
}

func initCache() {
	cacheTtl, err := strconv.Atoi(os.Getenv(ENV_CACHE_TTL_SECONDS))
	if err != nil {
		logWarning("Invalid CACHE_TTL_SECONDS environment variable, using default of 5 minutes.")
		cacheTtl = 300
	}
	ttl := time.Duration(cacheTtl) * time.Second

	appCache = &cache.Cache{
		Entries: make(map[string]cache.CacheEntry),
		Ttl:     ttl,
	}
}

func startServer() {
	// Try reading port from environment variable
	serverPort := os.Getenv(ENV_PORT)
	if serverPort == "" {
		logWarning("PORT environment variable not set, defaulting to 3000")
		serverPort = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/dashboard", handleDashboardData)

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: mux,
	}

	// Configure graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		logStartup("Server starting on port: " + serverPort)
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
