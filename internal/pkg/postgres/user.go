package postgres

import (
	"github.com/google/uuid"
	"github.com/nikovacevic/zyt/internal/app/zyt"
	"golang.org/x/crypto/bcrypt"
)

// SaveUser stores or updates the given User
func (db *DB) SaveUser(user *zyt.User) (*zyt.User, error) {
	if &user.ID == nil {
		// Hash password before storing new user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)

		insert := "INSERT INTO users (email, password) VALUES ($1, $2);"
		_, err = db.Exec(insert, user.Email, user.Password)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	// Update
	// TODO db.QueryRow("UPDATE user SET ...")

	return user, nil
}

// ViewUser retrieves the User with the given ID
func (db *DB) ViewUser(id uuid.UUID) (*zyt.User, error) {
	row := db.QueryRow("SELECT id, email FROM users WHERE id=$1;", id)

	user := zyt.User{}
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
