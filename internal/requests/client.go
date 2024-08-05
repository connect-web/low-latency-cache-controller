package requests

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// createHTTPClient creates an HTTP client with keep-alive connections enabled and sets a common cookie
func createHTTPClient(domain string, cookies map[string]string) *http.Client {
	// Create a cookie jar
	jar, _ := cookiejar.New(nil)

	// Create a custom transport with keep-alive settings
	transport := &http.Transport{
		DisableKeepAlives:     false,
		MaxIdleConns:          20,
		MaxIdleConnsPerHost:   20,
		IdleConnTimeout:       1 * time.Minute,
		TLSHandshakeTimeout:   20 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Jar:       jar, // Attach the cookie jar to the client
		Timeout:   60 * time.Second,
	}

	// Set a placeholder cookie to be used for all requests
	parsedURL, _ := url.Parse(domain)

	cookieJar := []*http.Cookie{}
	for cookieName, cookieValue := range cookies {
		cookieJar = append(cookieJar, &http.Cookie{
			Name:  cookieName,
			Value: cookieValue,
		})
	}
	jar.SetCookies(parsedURL, cookieJar)

	return client
}
