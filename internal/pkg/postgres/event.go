package postgres

import (
	"github.com/google/uuid"
	"github.com/nikovacevic/zyt/internal/app/zyt"
)

// SaveEvent stores or updates the given Event
func (db *DB) SaveEvent(event *zyt.Event) (*zyt.Event, error) {
	// TODO
	return nil, nil
}

// ViewEvent retrieves the Event with the given ID
func (db *DB) ViewEvent(id uuid.UUID) (*zyt.Event, error) {
	row := db.QueryRow("SELECT id, start_time, end_time, duration, name, notes FROM events WHERE id=$1;", id)

	event := zyt.Event{}
	err := row.Scan(&event.ID, &event.StartTime, &event.EndTime, &event.Duration, &event.Name, &event.Notes)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
