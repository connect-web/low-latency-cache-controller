package model

import (
	"fmt"
	"time"
)

type EntryHandler struct {
	Entries []ToplistEntry
}

type ToplistEntry struct {
	// Only the required fields are stored.
	Id          int       `json:"id"`
	LastUpdated time.Time `json:"lastUpdated"`
}

func NewEntryHandler() EntryHandler {
	return EntryHandler{Entries: make([]ToplistEntry, 0)}
}

func (handler *EntryHandler) FindOutdated(apiEntries EntryHandler) (invalidEntriesHandler EntryHandler) {
	// Finds rows that do not exist in api response or have mismatched update time.
	invalidEntriesHandler = NewEntryHandler()

	apiMap := apiEntries.getMap()

	for _, entry := range handler.Entries {
		apiTime, exists := apiMap[entry.Id]

		if !exists {
			fmt.Printf("%d does not exist in apiMap\n", entry.Id)
		}

		if !exists || !apiTime.Equal(entry.LastUpdated) {
			invalidEntriesHandler.Add(entry)
		}
	}
	return invalidEntriesHandler
}

func (handler *EntryHandler) Add(entry ToplistEntry) {
	handler.Entries = append(handler.Entries, entry)
}

func (handler *EntryHandler) getMap() map[int]time.Time {
	var timeMap = map[int]time.Time{}
	for _, entry := range handler.Entries {
		timeMap[entry.Id] = entry.LastUpdated
	}
	return timeMap
}

func (handler *EntryHandler) HasOutdatedEntries() bool {
	fmt.Printf("Found %d outdated entries!\n", len(handler.Entries))
	return 0 < len(handler.Entries)
}

func (handler *EntryHandler) BuildUrls(templateUrl string) []string {
	urls := []string{}
	for _, entry := range handler.Entries {
		urls = append(urls, fmt.Sprintf(templateUrl, entry.Id))
	}
	return urls
}
