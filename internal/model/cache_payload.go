package model

import "os"

type AuthCachePayload struct {
	CacheKey string   `json:"cacheKey"`
	Target   []string `json:"Target"`
}

func NewAuthCachePayload(urls []string) AuthCachePayload {
	return AuthCachePayload{
		CacheKey: os.Getenv("cacheKey"),
		Target:   urls,
	}
}
