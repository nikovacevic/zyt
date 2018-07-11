package postgres

import (
	"fmt"

	"github.com/nikovacevic/zyt/internal/app/zyt"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser authenticates a user's email and password combination.
func (db *DB) AuthenticateUser(email, password string) (*zyt.User, error) {
	var hashedPassword []byte
	row := db.QueryRow("SELECT id, email, password FROM users WHERE email=$1;", email)

	user := zyt.User{}
	err := row.Scan(&user.ID, &user.Email, &hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("User with email %s not found", email)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Incorrect username and password combination")
	}

	return &user, nil
}
