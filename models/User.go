package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

// User is a user
type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password []byte    `json:"password"`
}

// NewUser creates a new User instance
func NewUser(email string, password []byte) *User {
	return &User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}
}

// ToJSON marshals a User to JSON
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
