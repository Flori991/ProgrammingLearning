package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	initialize()
	startServer()
}

func initialize() {
	logStartup("Starting up airvpn-api.")

	// Initialize logger
	logStartup("Initializing Log Level.")
	initLogLevel()
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
