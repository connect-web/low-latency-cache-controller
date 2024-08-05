package ml

import (
	"fmt"
	"github.com/connect-web/low-latency-cache-controller/internal/db"
	"github.com/connect-web/low-latency-cache-controller/internal/requests"
)

func LoopUntilCacheUpdated(host string) {
	cached := compareTimestamps(host)
	for !cached {
		cached = compareTimestamps(host)
	}
}

// This method will reload the cache only if the pages are already cached and have incorrect timestamps
func compareTimestamps(host string) (finished bool) {
	database := db.NewDBClient()
	defer database.Close()

	response := requests.Cache(host, []string{"/api/v2/ml/skill-toplist", "/api/v2/ml/boss-minigame-toplist"}, false)
	apiSkillToplist, apiMinigameToplist := response.LoadMlToplists()

	dbSkillToplist, err := database.MlSkillToplist()
	if err != nil {
		fmt.Printf("Failed to fetch skill toplist: %s\n", err.Error())
		return false
	}

	dbMinigameToplist, err := database.MlMinigameToplist()
	if err != nil {
		fmt.Printf("Failed to fetch minigame toplist: %s\n", err.Error())
		return false
	}
	invalidSkillRows := dbSkillToplist.FindOutdated(apiSkillToplist)
	invalidMinigameRows := dbMinigameToplist.FindOutdated(apiMinigameToplist)

	if invalidSkillRows.HasOutdatedEntries() {
		skillUrls := invalidSkillRows.BuildUrls("/api/v2/ml/skill-toplist-users?skill=%s")
		requests.RefreshCache(host, skillUrls)
		requests.RefreshCache(host, []string{"/api/v2/ml/skill-toplist"})
	}

	if invalidMinigameRows.HasOutdatedEntries() {
		minigameUrls := invalidMinigameRows.BuildUrls("/api/v2/ml/boss-minigame-toplist-users?minigame=%s")
		requests.RefreshCache(host, minigameUrls)
		requests.RefreshCache(host, []string{"/api/v2/ml/boss-minigame-toplist"})
	}
	return !invalidSkillRows.HasOutdatedEntries() && !invalidMinigameRows.HasOutdatedEntries()
}
