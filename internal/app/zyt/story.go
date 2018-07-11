package zyt

import "github.com/google/uuid"

// Story is a sequnce of Events belonging to a User
type Story struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
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
