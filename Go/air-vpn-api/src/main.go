package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/Flori991/ProgrammingLearning/types"
)

func main() {
	// Fetch UserInfo
	userInfoResponseBytes := httpGet(AIRVPN_USERINFO_URL, ENV_API_KEY)
	if strings.Contains(string(userInfoResponseBytes), "Not authorized") {
		log.Fatal("API key is not valid. Please check your environment variable.")
	}

	//Fetch Server Status
	statusResponseBytes := httpGet(AIRVPN_STATUS_URL)

	// Turn userinfo json response into response type
	userInfo := safeJsonParse(userInfoResponseBytes, types.UserInfo{})
	if len(userInfo.Sessions) == 0 {
		log.Println("No active sessions found.")
	}

	// Turn status json response into response type
	status := safeJsonParse(statusResponseBytes, types.Status{})
	if len(status.Servers) == 0 {
		log.Println("No server status information found.")
	}

	// Merge responses into custom api struct
	sessionSummaries := mergeResponsesIntoSummaries(userInfo.Sessions, status.Servers)

	// Serialize for return
	jsonSessionSummaries, err := json.Marshal(sessionSummaries)
	if err != nil {
		log.Fatal("Something went wrong serializing the response:", err)
	}

	log.Println("API call successful. Response:", string(jsonSessionSummaries))
}
