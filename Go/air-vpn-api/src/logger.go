package main

import (
	"log"
)

func logStartup(message string) {
	log.Println(LOG_STARTUP, message)
}

func logDebug(message string) {
	if config.LogLevel >= LEVEL_DEBUG {
		log.Println(LOG_DEBUG, message)
	}
}

func logInfo(message string) {
	if config.LogLevel >= LEVEL_INFO {
		log.Println(LOG_INFO, message)
	}
}

func logWarning(message string) {
	if config.LogLevel >= LEVEL_WARNING {
		log.Println(LOG_WARNING, message)
	}
}

func logError(message string, err ...error) {
	if config.LogLevel >= LEVEL_ERROR {
		if len(err) > 0 {
			log.Println(LOG_ERROR, message, err[0])
		} else {
			log.Println(LOG_ERROR, message)
		}
	}
}
