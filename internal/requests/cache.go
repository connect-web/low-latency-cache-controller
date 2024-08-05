package requests

import (
	"github.com/connect-web/low-latency-cache-controller/internal/model"
	"os"
)

var (
	cookies = map[string]string{
		"session_id": os.Getenv("admin_session_id"),
	}
)

func RefreshCache(host string, urls []string) model.ResponseHandler {
	return Cache(host, urls, true)
}

func Cache(host string, urls []string, noCache bool) model.ResponseHandler {
	// Initialize the responseHandler object
	var responseHandler = model.ResponseHandler{}

	// Create a custom HTTP client
	client := createHTTPClient(host, cookies)
	defer client.CloseIdleConnections()

	// Iterate over the URLs and make requests
	for _, url := range urls {
		// Make the request and store the result
		result := cacheRequest(client, host+url, noCache)
		responseHandler.Results = append(responseHandler.Results, result)
	}
	// Prints the amount of successful , failed requests & mean connection time.
	responseHandler.DisplayStats()

	return responseHandler
}
