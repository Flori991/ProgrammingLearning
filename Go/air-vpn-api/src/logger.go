package main

import (
	"log"
)

func logStartup(message string) {
	log.Println(LogStartup, message)
}

func logDebug(message string) {
	if config.LogLevel >= LevelDebug {
		log.Println(LogDebug, message)
	}
}

func logInfo(message string) {
	if config.LogLevel >= LevelInfo {
		log.Println(LogInfo, message)
	}
}

func logWarning(message string) {
	if config.LogLevel >= LevelWarning {
		log.Println(LogWarning, message)
	}
}

func logError(message string, err ...error) {
	if config.LogLevel >= LevelError {
		if len(err) > 0 {
			log.Println(LogError, message, err[0])
		} else {
			log.Println(LogError, message)
		}
	}
}
