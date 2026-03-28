package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func httpGet(apiUrl string, keyEnvVar ...string) []byte {
	log.Println("Starting API call...")
	// Create the HTTP client
	client := &http.Client{}
	client.Timeout = time.Second * 3

	//Create the HTTP request
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Fatal("Something went wrong creating the request:", err)
	}

	// If API key specified, add it as header
	if len(keyEnvVar) > 0 {
		req.Header.Set("API-KEY", os.Getenv(keyEnvVar[0]))
	}

	// Actually do the request
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("The request failed:", err)
	}

	// Log the response status code
	// Read the response body and return it
	defer response.Body.Close()
	log.Println("Response status:", response.Status)
	body, err := io.ReadAll(response.Body)
	return body
}
