package requests

import (
	"github.com/connect-web/low-latency-cache-controller/internal/model"
	"sync"
)

// CacheFast uses 5 goroutines to cache URLs
func CacheFast(host string, urls []string) model.ResponseHandler {
	// Initialize the responseHandler object
	var responseHandler = model.ResponseHandler{Results: make([]model.Response, 0)}

	// Create a custom HTTP client
	client := createHTTPClient(host, cookies)
	defer client.CloseIdleConnections()

	// Channel to distribute URLs to workers
	urlChannel := make(chan string, len(urls))
	resultChannel := make(chan model.Response, len(urls))

	// Add all URLs to the channel
	for _, url := range urls {
		urlChannel <- host + url
	}
	close(urlChannel) // Close the channel after sending all URLs

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch 5 goroutines to process URLs
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urlChannel {
				result := cacheRequest(client, url, false)
				resultChannel <- result
			}
		}()
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Collect results
	for result := range resultChannel {
		responseHandler.Results = append(responseHandler.Results, result)
	}

	// Prints the amount of successful, failed requests & mean connection time.
	responseHandler.DisplayStats()

	return responseHandler
}
