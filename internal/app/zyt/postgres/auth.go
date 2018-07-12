package postgres

import (
	"fmt"

	"github.com/nikovacevic/zyt/internal/app/zyt"
)

// AuthenticateUser authenticates a user's email and password combination.
func (db *DB) AuthenticateUser(email, password string) (*zyt.User, error) {
	row := db.QueryRow(`
		SELECT   id, email FROM users
		WHERE    email = lower($1)
		AND      password = crypt($2, password);
	`, email, password)

	user := zyt.User{}
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("Incorrect username and password combination")
	}

	return &user, nil
}
