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

func (client *DBClient) MlSkillToplist() (entries model.MlEntryHandler, Error error) {
	entries = model.NewMlEntryHandler()

	query := `
	SELECT r.activity, max(r.Time)
	FROM ML.results r
	WHERE r.activitytype = 'skills'
	GROUP BY r.activity
	`

	rows, err := client.DB.Query(query)

	if err != nil {
		return entries, err
	}

	defer rows.Close()

	for rows.Next() {
		var entry model.MlToplistEntry

		scanErr := rows.Scan(&entry.Minigame, &entry.LastUpdated)

		if scanErr == nil {
			entries.Add(entry)
		}
	}

	return entries, nil
}

func (client *DBClient) MlMinigameToplist() (entries model.MlEntryHandler, Error error) {
	entries = model.NewMlEntryHandler()

	query := `
	SELECT r.activity, max(r.Time)
	FROM ML.results r
	WHERE r.activitytype = 'minigames'
	GROUP BY r.activity
	`
	rows, err := client.DB.Query(query)

	if err != nil {
		return entries, err
	}

	defer rows.Close()

	for rows.Next() {
		var entry model.MlToplistEntry

		scanErr := rows.Scan(&entry.Minigame, &entry.LastUpdated)

		if scanErr == nil {
			entries.Add(entry)
		}
	}

	return entries, nil
}
