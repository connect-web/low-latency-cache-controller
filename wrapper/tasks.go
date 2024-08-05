package wrapper

import (
	"github.com/connect-web/low-latency-cache-controller/internal/tasks"
	"github.com/connect-web/low-latency-cache-controller/internal/tasks/ml"
	"github.com/connect-web/low-latency-cache-controller/internal/tasks/public"
	"time"
)

func StartUp(host string) {
	time.Sleep(60 * time.Second)
	tasks.CacheAll(host)
}

func RefreshCacheHourly(host string) {
	time.Sleep(60 * time.Second)
	for {
		public.LoopUntilCacheUpdated(host)
		ml.LoopUntilCacheUpdated(host)
		time.Sleep(1 * time.Hour)
	}
}
