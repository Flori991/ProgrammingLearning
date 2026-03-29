package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Flori991/ProgrammingLearning/types"
)

func handleDashboardData(w http.ResponseWriter, r *http.Request) {
	logInfo("Received request for dashboard data. ")
	logInfo("Request received from: " + getAccessIp(r))

	// Get API key from request header
	API_KEY := r.Header.Get("API-KEY")
	if API_KEY == "" {
		logError("API key is missing")
		w.Header().Set("x-missing-field", "API-KEY")
		http.Error(w, "API key is missing", http.StatusBadRequest)
		return
	}

	// Fetch UserInfo
	userInfoResponseBytes, err := httpGet(AIRVPN_USERINFO_URL, API_KEY)
	if err != nil {
		logError("Failed to fetch user info:", err)
		http.Error(w, "Failed to fetch user info", http.StatusInternalServerError)
		return
	}
	if strings.Contains(string(userInfoResponseBytes), "Not authorized") {
		logWarning("API key is not valid")
		http.Error(w, "API key is not valid", http.StatusUnauthorized)
		return
	}

	//Fetch Server Status
	statusResponseBytes, err := httpGet(AIRVPN_STATUS_URL)
	if err != nil {
		logError("Failed to fetch server status:", err)
		http.Error(w, "Failed to fetch server status", http.StatusInternalServerError)
		return
	}

	// Turn userinfo json response into response type
	userInfo, err := safeJsonParse(userInfoResponseBytes, types.UserInfo{})
	if err != nil {
		logError("Failed to parse user info response:", err)
		http.Error(w, "Failed to parse user info response", http.StatusInternalServerError)
		return
	}
	if len(userInfo.Sessions) == 0 {
		logWarning("No active sessions found.")
		http.Error(w, "No active sessions found", http.StatusInternalServerError)
		return
	}

	// Turn status json response into response type
	status, err := safeJsonParse(statusResponseBytes, types.Status{})
	if err != nil {
		logError("Failed to parse server status response:", err)
		http.Error(w, "Failed to parse server status response", http.StatusInternalServerError)
		return
	}
	if len(status.Servers) == 0 {
		logWarning("No server status information found.")
		http.Error(w, "No server status information found", http.StatusInternalServerError)
		return
	}

	// Merge responses into custom api struct
	sessionSummaries := mergeResponsesIntoSummaries(userInfo.Sessions, status.Servers)

	// Serialize for return
	jsonSessionSummaries, err := json.Marshal(sessionSummaries)
	if err != nil {
		logError("Failed to serialize response:", err)
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.Writer.Write(w, jsonSessionSummaries)
}
