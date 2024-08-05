package model

import (
	"encoding/json"
	"fmt"
	"github.com/montanaflynn/stats"
	"strings"
)

type ResponseHandler struct {
	Results []Response
}

type Response struct {
	Url        string
	Valid      bool
	Response   []byte
	DurationMs float64
}

func (handler *ResponseHandler) LoadToplists() (skillToplist EntryHandler, minigameToplist EntryHandler) {
	skillToplist = NewEntryHandler()
	minigameToplist = NewEntryHandler()

	for _, result := range handler.Results {
		if strings.Contains(result.Url, "v2/public/skill-toplist") {
			err := json.Unmarshal(result.Response, &skillToplist.Entries)
			if err != nil {
				fmt.Println("Failed to decode skill toplist into struct.")
			}
		}
		if strings.Contains(result.Url, "v2/public/boss-minigame-toplist") {
			err := json.Unmarshal(result.Response, &minigameToplist.Entries)
			if err != nil {
				fmt.Println("Failed to decode minigame toplist into struct.")
			}
		}
	}
	return skillToplist, minigameToplist
}

func (handler *ResponseHandler) LoadMlToplists() (skillToplist MlEntryHandler, minigameToplist MlEntryHandler) {
	skillToplist = NewMlEntryHandler()
	minigameToplist = NewMlEntryHandler()

	for _, result := range handler.Results {
		if strings.Contains(result.Url, "api/v2/ml/skill-toplist") {
			err := json.Unmarshal(result.Response, &skillToplist.Entries)
			if err != nil {
				fmt.Println("Failed to decode ML skill toplist into struct.")
			}
		}
		if strings.Contains(result.Url, "v2/ml/boss-minigame-toplist") {
			err := json.Unmarshal(result.Response, &minigameToplist.Entries)
			if err != nil {
				fmt.Println("Failed to decode ML minigame toplist into struct.")
			}
		}
	}
	return skillToplist, minigameToplist
}

func (handler *ResponseHandler) DisplayStats() {
	validRequests, invalidRequests, meanRequestTime := handler.getStats()
	fmt.Printf("%d/%d valid requests, %d were invalid\nAverage Response MS: %.2f\n",
		validRequests, len(handler.Results), invalidRequests, meanRequestTime)
}

func (handler *ResponseHandler) getStats() (int, int, float64) {
	var validRequests int
	var invalidRequests int
	var durations []float64
	var meanDuration float64

	for _, result := range handler.Results {
		if result.Valid {
			validRequests++
			durations = append(durations, result.DurationMs)
		} else {
			invalidRequests++
		}
	}
	if 0 < len(durations) {
		meanDuration, _ = stats.Mean(durations)
	} else {
		fmt.Println("Cannot calculate mean duration for 0 successful requests.")
	}

	return validRequests, invalidRequests, meanDuration
}
