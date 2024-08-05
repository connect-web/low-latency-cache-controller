package main

import (
	"github.com/connect-web/low-latency-cache-controller/internal/tasks"
	"github.com/connect-web/low-latency-cache-controller/internal/tasks/ml"
	"github.com/connect-web/low-latency-cache-controller/internal/tasks/public"
)

var (
	onlineDomain = "https://low-latency.co.uk"
	localhost    = "http://127.0.0.1:4050"
)

func main() {
	tasks.CacheAll(localhost)
	public.LoopUntilCacheUpdated(localhost)
	ml.LoopUntilCacheUpdated(localhost)
}
