package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
)

func httpGet(apiUrl string, apiKey ...string) ([]byte, error) {
	requestName := path.Base(apiUrl)

	// Checking cache
	cacheKey := "dashboard:" + requestName
	// Add user specific if possible, to allow caching of user specific data
	if len(apiKey) > 0 {
		cacheKey += "/" + fmt.Sprintf("%x", sha256.Sum256([]byte(apiKey[0])))
	}
	logDebug("Checking cache for: " + requestName)
	if cachedBody, ok := appCache.Get(cacheKey); ok {
		logDebug("Found cache entry, returning it.")
		return cachedBody, nil
	}

	logDebug("Starting API call to: " + requestName)
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
	startTime := time.Now()
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	logDebug("API call for " + requestName + " completed in " + time.Since(startTime).String())

	// Log the response status code
	// Read the response body and return it
	defer response.Body.Close()
	logDebug("Response status for " + requestName + ": " + response.Status)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Set cache
	logDebug("No cache found, setting cache for " + requestName)
	appCache.Set(cacheKey, body)

	return body, nil
}
