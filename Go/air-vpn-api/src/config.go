package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Port            string
	CacheTtlSeconds time.Duration
	LogLevel        int
}

func initConfig() {
	logStartup("Initializing configuration.")
	config = &Config{
		Port:            parsePort(),
		CacheTtlSeconds: parseCacheTTL(),
		LogLevel:        parseLogLevel(),
	}
}

func parsePort() string {
	port := os.Getenv(ENV_PORT)
	if _, err := strconv.Atoi(port); err != nil || port == "" {
		logStartup("Invalid PORT, defaulting to 3000")
		return "3000"
	}
	return port
}

func parseCacheTTL() time.Duration {
	seconds, err := strconv.Atoi(os.Getenv(ENV_CACHE_TTL_SECONDS))
	if err != nil {
		logStartup("Invalid CACHE_TTL_SECONDS, defaulting to 300 seconds")
		return 300 * time.Second
	}
	return time.Duration(seconds) * time.Second
}

func parseLogLevel() int {
	switch strings.ToLower(os.Getenv(ENV_LOG_LEVEL)) {
	case "error", "0":
		return LEVEL_ERROR
	case "info", "2":
		return LEVEL_INFO
	case "debug", "3":
		return LEVEL_DEBUG
	default:
		logStartup("Invalid LOG_LEVEL, defaulting to WARNING")
		return LEVEL_WARNING
	}
}
