package store

import (
	"github.com/google/uuid"
	"github.com/nikovacevic/zyt-api/model"
)

// Event provides operations for persistent storage of Events et al.
type Event interface {
	SaveEvent(event *model.Event) error
	ViewEvent(id uuid.UUID) (*model.Event, error)
}

// SaveEvent stores or updates the given Event
func (db *DB) SaveEvent(event *model.Event) error {
	// TODO
	return nil
}

// ViewEvent retrieves the Event with the given ID
func (db *DB) ViewEvent(id uuid.UUID) (*model.Event, error) {
	// TODO
	return nil, nil
}
