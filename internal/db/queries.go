package db

import (
	"github.com/connect-web/low-latency-cache-controller/internal/model"
)

func (client *DBClient) SkillToplist() (entries model.EntryHandler, Error error) {
	entries = model.NewEntryHandler()

	query := "SELECT id, updated FROM grouped.skillers"

	rows, err := client.DB.Query(query)

	if err != nil {
		return entries, err
	}

	defer rows.Close()

	for rows.Next() {
		var entry model.ToplistEntry

		scanErr := rows.Scan(&entry.Id, &entry.LastUpdated)

		if scanErr == nil {
			entries.Add(entry)
		}
	}

	return entries, nil
}

func (client *DBClient) MinigameToplist() (entries model.EntryHandler, Error error) {
	entries = model.NewEntryHandler()

	query := "SELECT id, updated FROM grouped.minigames"

	rows, err := client.DB.Query(query)

	if err != nil {
		return entries, err
	}

	defer rows.Close()

	for rows.Next() {
		var entry model.ToplistEntry

		scanErr := rows.Scan(&entry.Id, &entry.LastUpdated)

		if scanErr == nil {
			entries.Add(entry)
		}
	}

	return entries, nil
}
