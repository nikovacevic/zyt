package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nikovacevic/zyt-api/model"
)

// User provides operations for persistent storage of Users et al.
type User interface {
	SaveUser(event *model.User) error
	ViewUser(id uuid.UUID) (*model.User, error)
}

// SaveUser stores or updates the given User
func (db *DB) SaveUser(event *model.User) error {
	// TODO
	return nil
}

// ViewUser retrieves the User with the given ID
func (db *DB) ViewUser(id uuid.UUID) (*model.User, error) {
	row := db.QueryRow("SELECT id, email FROM users WHERE id = $1", id)

	user := model.User{}
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("User not found")
	}

	return &user, nil
}
