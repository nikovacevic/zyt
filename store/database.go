package store

import (
	"database/sql"

	_ "github.com/lib/pq" // Postgres driver
)

// DB is a database handle
type DB struct {
	*sql.DB
}

// NewDB connects to a database, tests the connection, and returns a connection pool
func NewDB(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
