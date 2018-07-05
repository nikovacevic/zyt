package models

import (
	"encoding/json"
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

// Ending configures an event to end at the given time
func Ending(end time.Time) func(e *Event) {
	return func(e *Event) {
		e.EndTime = end

		if &e.StartTime != nil {
			e.Duration = e.EndTime.Sub(e.StartTime)
		}
	}
}

// ForUser configures an event to belong to the given User
func ForUser(user User) func(e *Event) {
	return func(e *Event) {
		e.UserID = user.ID
	}
}

// ForStory configures an event to belong to the given Story
func ForStory(story Story) func(e *Event) {
	return func(e *Event) {
		e.StoryID = story.ID
		e.UserID = story.UserID
	}
}

// NewEvent creates a new Event instance
func NewEvent(configs ...func(*Event)) *Event {
	e := &Event{ID: uuid.New()}

	for _, configure := range configs {
		configure(e)
	}

	return e
}

// ToJSON marshals a Event to JSON
func (e *Event) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}
