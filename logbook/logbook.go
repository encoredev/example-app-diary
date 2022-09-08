package logbook

import (
	"context"
	"encore.dev/storage/sqldb"
	"fmt"
	"time"
)

type LogbookEntry struct {
	Id   int
	Time time.Time
	Text string
}

type ListResponse struct {
	Entries []LogbookEntry
}

//encore:api public method=GET path=/logbook/entries/:date
func ListEntries(ctx context.Context, date string) (*ListResponse, error) {
	rows, err := sqldb.Query(ctx, `SELECT id, created_at, text FROM logbook WHERE date = $1`, date)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entries []LogbookEntry
	for rows.Next() {
		var entry = LogbookEntry{}
		if err := rows.Scan(&entry.Id, &entry.Time, &entry.Text); err != nil {
			return nil, fmt.Errorf("could not scan: %v", err)
		}
		entries = append(entries, entry)
	}

	return &ListResponse{Entries: entries}, nil
}

type CreateParams struct {
	Text string
}

//encore:api public method=POST path=/logbook/entries
func CreateEntry(ctx context.Context, params *CreateParams) (*LogbookEntry, error) {
	entry := &LogbookEntry{}
	err := sqldb.QueryRow(ctx, `
		INSERT INTO logbook (date, created_at, text)
		VALUES (CURRENT_DATE, NOW(), $1)
		RETURNING id, created_at, text
	`, params.Text).Scan(&entry.Id, &entry.Time, &entry.Text)

	if err != nil {
		return nil, err
	}

	return entry, nil
}
