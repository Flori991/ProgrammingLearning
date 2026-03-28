package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	types "github.com/Flori991/ProgrammingLearning/types"
)

const API_URL = "https://airvpn.org/api/userinfo/?key="
const ENV_API_KEY = "AIRVPN_API_KEY"

func main() {
	// Do the request and return json response
	responsebytes := sendGetRequest()
	if strings.Contains(string(responsebytes), "Not authorized") {
		log.Fatal("API key is not valid. Please check your environment variable.")
	}
	// Turn json response intro response type
	userInfo := parseResponse(responsebytes)
	if len(userInfo.Sessions) == 0 {
		log.Println("No active sessions found.")
	}
	// Turn response into custom struct and serialize it for return
	json := serializeResponse(userInfo.Sessions)

	log.Println("API call successful. Response:", json)
}

func sendGetRequest() []byte {
	log.Println("Starting API call...")
	resp, err := http.Get(API_URL + os.Getenv(ENV_API_KEY))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	return body
}

func parseResponse(body []byte) types.UserInfo {
	var userInfo types.UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		log.Fatal(err)
	}
	return userInfo
}

func mapToSessionSummary(session types.Session) types.SessionSummary {
	return types.SessionSummary{
		DeviceName:         session.DeviceName,
		DeviceDescription:  session.DeviceDescription,
		ConnectedSinceDate: session.ConnectedSinceDate,
		ExitIp:             session.ExitIp,
	}
}

func serializeResponse(response []types.Session) string {
	json, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
