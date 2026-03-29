package main

import (
	"log"
	"os"
	"strings"
)

var logLevel = LEVEL_WARNING

func logStartup(message string) {
	log.Println(LOG_STARTUP, message)
}

func initLogLevel() {
	level := os.Getenv(ENV_LOG_LEVEL)
	switch strings.ToLower(level) {
	case "error", "0":
		logLevel = LEVEL_ERROR
	case "warning", "1":
		logLevel = LEVEL_WARNING
	case "info", "2":
		logLevel = LEVEL_INFO
	case "debug", "3":
		logLevel = LEVEL_DEBUG
	default:
		logLevel = LEVEL_WARNING
		logWarning("Invalid LOG_LEVEL environment variable, defaulting to WARNING")
	}
}

func logDebug(message string) {
	if logLevel >= LEVEL_DEBUG {
		log.Println(LOG_DEBUG, message)
	}
}

func logInfo(message string) {
	if logLevel >= LEVEL_INFO {
		log.Println(LOG_INFO, message)
	}
}

func logWarning(message string) {
	if logLevel >= LEVEL_WARNING {
		log.Println(LOG_WARNING, message)
	}
}

func logError(message string, err ...error) {
	if logLevel >= LEVEL_ERROR {
		if len(err) > 0 {
			log.Println(LOG_ERROR, message, err[0])
		} else {
			log.Println(LOG_ERROR, message)
		}
	}
}
