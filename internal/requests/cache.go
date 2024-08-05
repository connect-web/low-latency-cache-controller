package requests

import (
	"fmt"
	"github.com/connect-web/low-latency-cache-controller/internal/model"
	"os"
)

var (
	cookies = map[string]string{
		"session_id": os.Getenv("admin_session_id"),
	}
)

func Cache(host string, urls []string) model.ResponseHandler {
	// Initialize the responseHandler object
	var responseHandler = model.ResponseHandler{}

	// Create a custom HTTP client
	client := createHTTPClient(host, cookies)
	defer client.CloseIdleConnections()

	// Iterate over the URLs and make requests
	for _, url := range urls {
		// Make the request and store the result
		result := cacheRequest(client, host+url)
		responseHandler.Results = append(responseHandler.Results, result)
	}
	// Prints the amount of successful , failed requests & mean connection time.
	responseHandler.DisplayStats()

	return responseHandler
}

func RefreshCache(host string, urls []string) model.ResponseHandler {
	// Initialize the responseHandler object
	var responseHandler = model.ResponseHandler{}

	// Create a custom HTTP client
	client := createHTTPClient(host, cookies)
	defer client.CloseIdleConnections()

	// Iterate over the URLs and make requests
	for _, urlBatch := range chunkUrls(urls, 20) {
		// Make the request and store the result
		valid, err := revokeCacheRequest(client, host, urlBatch)
		if !valid {
			fmt.Printf("Failed to revoke cache for %d urls due to: %s\n", len(urls), err.Error())
			continue
		}
		// cache has been deleted for urlBatch, now it needs re-caching
		Cache(host, urlBatch)
	}

	return responseHandler
}

func chunkUrls(urls []string, size int) [][]string {
	var chunks = make([][]string, 0)
	var temporary = make([]string, 0)

	for _, url := range urls {
		temporary = append(temporary, url)
		if len(temporary)%size == 0 {
			chunks = append(chunks, temporary)
			temporary = make([]string, 0)
		}
	}
	if 0 < len(temporary) {
		chunks = append(chunks, temporary)
		temporary = make([]string, 0)
	}
	return chunks
}
