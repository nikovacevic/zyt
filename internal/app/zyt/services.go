package zyt

import (
	"github.com/google/uuid"
)

// AuthService provides authentication and authorization features
type AuthService interface {
	AuthenticateUser(email, password string) (*User, error)
}

// ErrorService provides error-related features and utilities
type ErrorService interface {
	CheckAndFatal(error)
	CheckAndLog(error)
	CheckAndPanic(error)
}

// EventService provides features for viewing and transforming events
type EventService interface {
	// TODO EndEvent(event *Event) (*Event, error)
	SaveEvent(event *Event) (*Event, error)
	// TODO StartEvent(configs ...func(*Event)) (*Event, error)
	ViewEvent(id uuid.UUID) (*Event, error)
}

// TokenService provides API tokens
type TokenService interface {
	GenerateToken(user *User) (string, error)
	VerifyToken(tokenString string) error
}

// UserService provides features for viewing and transforming users
type UserService interface {
	SaveUser(user *User) (*User, error)
	ViewUser(id uuid.UUID) (*User, error)
}
