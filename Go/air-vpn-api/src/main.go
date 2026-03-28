package main

import (
	"encoding/json"
	"log"
	"strings"

	types "github.com/Flori991/ProgrammingLearning/types"
)

func main() {
	// Fetch UserInfo
	userInfoResponseBytes := sendGetRequest(AIRVPN_USERINFO_URL, ENV_API_KEY)
	if strings.Contains(string(userInfoResponseBytes), "Not authorized") {
		log.Fatal("API key is not valid. Please check your environment variable.")
	}

	//Fetch Server Status
	statusResponseBytes := sendGetRequest(AIRVPN_STATUS_URL)
	log.Println("Server status response:", string(statusResponseBytes))

	// Turn json response intro response type
	userInfo := parseResponse(userInfoResponseBytes)
	if len(userInfo.Sessions) == 0 {
		log.Println("No active sessions found.")
	}
	// Turn response into custom struct and serialize it for return
	jsonSessionSummaries := serializeResponse(userInfo.Sessions)

	log.Println("API call successful. Response:", jsonSessionSummaries)
}

func parseResponse(body []byte) types.UserInfo {
	var userInfo types.UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		log.Fatal(err)
	}
	return userInfo
}

func mapToSessionSummaries(sessions []types.Session) types.SessionSummaries {
	var sessionSummaries []types.SessionSummary
	for _, session := range sessions {
		sessionSummaries = append(sessionSummaries, types.SessionSummary{
			DeviceName:         session.DeviceName,
			DeviceDescription:  session.DeviceDescription,
			ExitIpv4:           session.ExitIpv4,
			ServerName:         session.ServerName,
			ServerCountry:      session.ServerCountry,
			BytesRead:          session.BytesRead,
			BytesWrite:         session.BytesWrite,
			ConnectedSinceDate: session.ConnectedSinceDate,
			ConnectedSinceUnix: session.ConnectedSinceUnix,
		})
	}
	return types.SessionSummaries{
		Sessions: sessionSummaries,
	}
}

func serializeResponse(sessions []types.Session) string {
	json, err := json.Marshal(mapToSessionSummaries(sessions))
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
