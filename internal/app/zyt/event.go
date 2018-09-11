package zyt

import (
	"time"

	"github.com/google/uuid"
)

// Event is an event belonging to a User
type Event struct {
	ID        uuid.UUID     `json:"id"`
	StoryID   uuid.UUID     `json:"story_id"`
	UserID    uuid.UUID     `json:"user_id"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	Duration  time.Duration `json:"duration"`
	Name      string        `json:"name"`
	Notes     string        `json:"notes"`
}

// NewEvent creates a new Event instance
func NewEvent(configs ...func(*Event)) *Event {
	e := &Event{ID: uuid.New()}

	for _, configure := range configs {
		configure(e)
	}

	return e
}

// Starting configures an event to start at the given time
func Starting(start time.Time) func(e *Event) {
	return func(e *Event) {
		e.StartTime = start

		if &e.EndTime != nil {
			e.Duration = e.EndTime.Sub(e.StartTime)
		}
	}
}
