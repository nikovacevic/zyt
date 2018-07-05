package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Story is a sequnce of Events belonging to a User
type Story struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
}

// WithName configures a Story with a name
func WithName(name string) func(s *Story) {
	return func(s *Story) {
		s.Name = name
	}
}

// NewStory creates a new Story instance with configuration
func NewStory(userID uuid.UUID, configs ...func(*Story)) *Story {
	s := &Story{
		ID:     uuid.New(),
		UserID: userID,
	}

	for _, configure := range configs {
		configure(s)
	}

	return s
}

// ToJSON marshals a Story to JSON
func (s *Story) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
