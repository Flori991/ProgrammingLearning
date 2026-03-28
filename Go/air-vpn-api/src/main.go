package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Flori991/ProgrammingLearning/types"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/dashboard", handleDashboardData)

	log.Println("Starting server on port: " + SERVER_PORT)
	err := http.ListenAndServe(":"+SERVER_PORT, mux)
	if err != nil {
		log.Println("Server failed to start:", err)
	}
}

func handleDashboardData(w http.ResponseWriter, r *http.Request) {
	// Get API key from request header
	API_KEY := r.Header.Get("API-KEY")
	if API_KEY == "" {
		w.Header().Set("x-missing-field", "API-KEY")
		http.Error(w, "API key is missing", http.StatusBadRequest)
		return
	}
	log.Println(API_KEY)

	// Fetch UserInfo
	userInfoResponseBytes, err := httpGet(AIRVPN_USERINFO_URL, API_KEY)
	if err != nil {
		log.Println("Failed to fetch user info:", err)
		http.Error(w, "Failed to fetch user info", http.StatusInternalServerError)
		return
	}
	if strings.Contains(string(userInfoResponseBytes), "Not authorized") {
		http.Error(w, "API key is not valid", http.StatusUnauthorized)
		return
	}

	//Fetch Server Status
	statusResponseBytes, err := httpGet(AIRVPN_STATUS_URL)
	if err != nil {
		log.Println("Failed to fetch server status:", err)
		http.Error(w, "Failed to fetch server status", http.StatusInternalServerError)
		return
	}

	// Turn userinfo json response into response type
	userInfo, err := safeJsonParse(userInfoResponseBytes, types.UserInfo{})
	if err != nil {
		log.Println("Failed to parse user info response:", err)
		http.Error(w, "Failed to parse user info response", http.StatusInternalServerError)
		return
	}
	if len(userInfo.Sessions) == 0 {
		log.Println("No active sessions found.")
		http.Error(w, "No active sessions found", http.StatusInternalServerError)
		return
	}

	// Turn status json response into response type
	status, err := safeJsonParse(statusResponseBytes, types.Status{})
	if err != nil {
		log.Println("Failed to parse server status response:", err)
		http.Error(w, "Failed to parse server status response", http.StatusInternalServerError)
		return
	}
	if len(status.Servers) == 0 {
		log.Println("No server status information found.")
		http.Error(w, "No server status information found", http.StatusInternalServerError)
		return
	}

	// Merge responses into custom api struct
	sessionSummaries := mergeResponsesIntoSummaries(userInfo.Sessions, status.Servers)

	// Serialize for return
	jsonSessionSummaries, err := json.Marshal(sessionSummaries)
	if err != nil {
		log.Println("Failed to serialize response:", err)
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	io.Writer.Write(w, jsonSessionSummaries)
}
