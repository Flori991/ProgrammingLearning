package main

const AIRVPN_USERINFO_URL = "https://airvpn.org/api/userinfo/?key="
const AIRVPN_STATUS_URL = "https://airvpn.org/api/status/"
const ENV_LOG_LEVEL = "LOG_LEVEL"
const ENV_PORT = "PORT"

const LOG_ERROR = "[ERROR] "
const LOG_WARNING = "[WARNING] "
const LOG_INFO = "[INFO] "
const LOG_DEBUG = "[DEBUG] "
const LOG_STARTUP = "[STARTUP] "

const (
	LEVEL_ERROR   = 0
	LEVEL_WARNING = 1
	LEVEL_INFO    = 2
	LEVEL_DEBUG   = 3
)
