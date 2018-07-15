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
	// TODO
	return nil, nil
}
