package tasks

import (
	"fmt"
	"github.com/connect-web/low-latency-cache-controller/internal/requests"
)

// This will cache all pages
func CacheAll(host string) {
	// Cache Public API's
	cacheAllPublic(host)

	// Cache ML API's
	cacheAllMl(host)

}

func cacheAllPublic(host string) {
	response := requests.Cache(host, []string{"/api/v2/public/skill-toplist", "/api/v2/public/boss-minigame-toplist"}, false)
	apiSkillToplist, apiMinigameToplist := response.LoadToplists()
	skillUrls := apiSkillToplist.BuildUrls("/api/v2/public/skill-toplist-users?skill-id=%d")
	minigameUrls := apiMinigameToplist.BuildUrls("/api/v2/public/boss-minigame-toplist-users?minigame-id=%d")
	fmt.Printf("Caching %d public urls.\n", len(skillUrls)+len(minigameUrls))
	requests.CacheFast(host, skillUrls)
	requests.CacheFast(host, minigameUrls)
}

func cacheAllMl(host string) {
	response := requests.Cache(host, []string{"/api/v2/ml/skill-toplist", "/api/v2/ml/boss-minigame-toplist"}, false)
	apiSkillToplist, apiMinigameToplist := response.LoadMlToplists()
	skillUrls := apiSkillToplist.BuildUrls("/api/v2/ml/skill-toplist-users?skill=%s")
	minigameUrls := apiMinigameToplist.BuildUrls("/api/v2/ml/boss-minigame-toplist-users?minigame=%s")
	fmt.Printf("Caching %d ML urls.\n", len(skillUrls)+len(minigameUrls))
	requests.CacheFast(host, skillUrls)
	requests.CacheFast(host, minigameUrls)
}
