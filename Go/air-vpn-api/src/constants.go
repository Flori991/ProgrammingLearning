package main

const AirVPNUserInfoURL = "https://airvpn.org/api/userinfo/"
const AirVPNStatusURL = "https://airvpn.org/api/status/"
const EnvLogLevel = "LOG_LEVEL"
const EnvPort = "PORT"
const EnvCacheTTLSeconds = "CACHE_TTL_SECONDS"

const LogError = "[ERROR] "
const LogWarning = "[WARNING] "
const LogInfo = "[INFO] "
const LogDebug = "[DEBUG] "
const LogStartup = "[STARTUP] "

const (
	LevelError   = 0
	LevelWarning = 1
	LevelInfo    = 2
	LevelDebug   = 3
)
