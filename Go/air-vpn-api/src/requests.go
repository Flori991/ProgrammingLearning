package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func httpGet(apiUrl string, apiKey ...string) ([]byte, error) {
	log.Println("Starting API call...")
	// Create the HTTP client
	client := &http.Client{}
	client.Timeout = time.Second * 3

	//Create the HTTP request
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	// If API key specified, add it as header
	if len(apiKey) > 0 {
		req.Header.Set("API-KEY", apiKey[0])
	}

	// Actually do the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Log the response status code
	// Read the response body and return it
	defer response.Body.Close()
	log.Println("Response status:", response.Status)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
