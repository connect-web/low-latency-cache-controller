package requests

import (
	"fmt"
	"github.com/connect-web/low-latency-cache-controller/internal/model"
	"io"
	"log"
	"net/http"
	"time"
)

var cacheRequests = 0

// makeRequest performs an HTTP GET request to the specified URL
func cacheRequest(client *http.Client, url string, noCache bool) model.Response {
	// Create the result object
	var result model.Response
	result.Url = url

	// start a timer
	requestInit := time.Now()

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Failed to create request for %s: %v\n", url, err)
		return result
	}

	if noCache {
		req.Header.Set("Cache-Control", "no-cache")
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to perform request to %s: %v\n", url, err)
		return result
	}
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		//  Error reading body
		log.Printf("Error reading body: %s\n", readErr.Error())
		return result
	}

	// Store response
	result.Response = bodyBytes

	// Calculate duration.
	result.DurationMs = float64(time.Now().UnixMilli() - requestInit.UnixMilli())

	// set Valid = True If the request was successful
	result.Valid = resp.StatusCode == http.StatusOK && readErr == nil

	if result.Valid {
		cacheRequests++
		if cacheRequests%10 == 0 {
			fmt.Printf("%d pages cached.\n", cacheRequests)
		}
	}

	return result
}
