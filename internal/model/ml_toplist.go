package model

import (
	"fmt"
	"time"
)

type MlEntryHandler struct {
	Entries []MlToplistEntry
}

type MlToplistEntry struct {
	Minigame    string    `json:"minigame"`
	LastUpdated time.Time `json:"lastUpdated"`
}

func NewMlEntryHandler() MlEntryHandler {
	return MlEntryHandler{Entries: make([]MlToplistEntry, 0)}
}

func (handler *MlEntryHandler) FindOutdated(apiEntries MlEntryHandler) (invalidEntriesHandler MlEntryHandler) {
	// Finds rows that do not exist in api response or have mismatched update time.
	invalidEntriesHandler = NewMlEntryHandler()

	apiMap := apiEntries.getMap()

	for _, entry := range handler.Entries {
		apiTime, exists := apiMap[entry.Minigame]

		if !exists {
			fmt.Printf("%d does not exist in apiMap\n", entry.Minigame)
		}

		if !exists || !apiTime.Equal(entry.LastUpdated) {
			fmt.Printf("%v != %v [%s]\n", entry.LastUpdated, apiTime, entry.Minigame)
			invalidEntriesHandler.Add(entry)
		}
	}
	return invalidEntriesHandler
}

func (handler *MlEntryHandler) Add(entry MlToplistEntry) {
	handler.Entries = append(handler.Entries, entry)
}

func (handler *MlEntryHandler) getMap() map[string]time.Time {
	var timeMap = map[string]time.Time{}
	for _, entry := range handler.Entries {
		timeMap[entry.Minigame] = entry.LastUpdated
	}
	return timeMap
}

func (handler *MlEntryHandler) HasOutdatedEntries() bool {
	fmt.Printf("[ML] Found %d outdated entries!\n", len(handler.Entries))
	return 0 < len(handler.Entries)
}

func (handler *MlEntryHandler) BuildUrls(templateUrl string) []string {
	urls := []string{}
	for _, entry := range handler.Entries {
		urls = append(urls, fmt.Sprintf(templateUrl, entry.Minigame))
	}
	return urls
}
