package zyt

import (
	"time"

	"github.com/google/uuid"
)

// Session is a user
type Session struct {
	Token     uuid.UUID `json:"token"`
	createdAt time.Time
	expiresAt time.Time
}

// NewSession creates a new Session instance
func NewSession(user *User) *Session {
	tok := uuid.New()
	now := time.Now()
	exp := now.Add(100 * 24 * time.Hour)

	return &Session{
		Token:     tok,
		createdAt: now,
		expiresAt: exp,
	}
}

// IsExpired returns true if a session is currently expired and false if not.
func (s *Session) IsExpired() bool {
	return s.expiresAt.Before(time.Now())
}
