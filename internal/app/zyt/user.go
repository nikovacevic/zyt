package zyt

import "github.com/google/uuid"

// User is a user
type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// NewUser creates a new User instance
func NewUser(email, password string) *User {
	return &User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}
}
