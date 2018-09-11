package zyt

import (
	"github.com/google/uuid"
)

// AuthService provides authentication and authorization features
type AuthService interface {
	SessionService
	AuthenticateUser(email, password string) (*User, error)
}

// ErrorService provides error-related features and utilities
type ErrorService interface {
	CheckFatal(error)
	CheckLog(error)
	CheckPanic(error)
}

// EventService provides features for viewing and transforming events
type EventService interface {
	// TODO EndEvent(event *Event) (*Event, error)
	SaveEvent(event *Event) (*Event, error)
	// TODO StartEvent(configs ...func(*Event)) (*Event, error)
	ViewEvent(id uuid.UUID) (*Event, error)
}

// SessionService provides, maintains, and revokes API access using tokens
type SessionService interface {
	CreateSession(user *User) (*Session, error)
	VerifySession(token []byte) (bool, error)
	RevokeSession(token []byte) error
}

// UserService provides features for viewing and transforming users
type UserService interface {
	SaveUser(user *User) (*User, error)
	ViewUser(id uuid.UUID) (*User, error)
}
