package zyt

import "github.com/google/uuid"

// AuthService provides authentication and authorization features
type AuthService interface {
	AuthenticateUser(email, password string) (*User, error)
}

// EventService provides features for viewing and transforming events
type EventService interface {
	// EndEvent(event *Event) (*Event, error)
	SaveEvent(user *Event) (*Event, error)
	// StartEvent(configs ...func(*Event)) (*Event, error)
	ViewEvent(id uuid.UUID) (*Event, error)
}

// UserService provides features for viewing and transforming users
type UserService interface {
	SaveUser(user *User) (*User, error)
	ViewUser(id uuid.UUID) (*User, error)
}
