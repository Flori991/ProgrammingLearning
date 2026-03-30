package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Port     string
	CacheTtl time.Duration
	LogLevel int
}

func initConfig() {
	logStartup("Initializing configuration.")
	config = &Config{
		Port:     parsePort(),
		CacheTtl: parseCacheTTL(),
		LogLevel: parseLogLevel(),
	}
}

func parsePort() string {
	port := os.Getenv(EnvPort)
	if _, err := strconv.Atoi(port); err != nil || port == "" {
		logStartup("Invalid PORT, defaulting to 3000")
		return "3000"
	}
	return port
}

func parseCacheTTL() time.Duration {
	seconds, err := strconv.Atoi(os.Getenv(EnvCacheTTLSeconds))
	if err != nil {
		logStartup("Invalid CACHE_TTL_SECONDS, defaulting to 300 seconds")
		return 300 * time.Second
	}
	return time.Duration(seconds) * time.Second
}

func parseLogLevel() int {
	switch strings.ToLower(os.Getenv(EnvLogLevel)) {
	case "error", "0":
		return LevelError
	case "warning", "1":
		return LevelWarning
	case "info", "2":
		return LevelInfo
	case "debug", "3":
		return LevelDebug
	default:
		logStartup("Invalid LOG_LEVEL, defaulting to WARNING")
		return LevelWarning
	}
}
