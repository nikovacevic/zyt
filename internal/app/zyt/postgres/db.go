package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/spf13/viper"
)

// Config stores configuration setting for connecting to Postgres DB
type Config struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	Params   string
}

// DB is a database handle
type DB struct {
	*sql.DB
}

// GetConfig finds a config file in the given path and constructs a Config
func GetConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cnf Config
	if err := v.Unmarshal(&cnf); err != nil {
		return nil, err
	}

	return &cnf, nil
}

// DataSourceName constructs a dataSourceName (for opening connections to a SQL
// DB) from a Config
func (cnf Config) DataSourceName() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s",
		cnf.Username,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
		cnf.Params,
	)
}

// NewDB connects to a database, tests the connection, and returns a connection pool
func NewDB(cnf Config) (*DB, error) {
	db, err := sql.Open("postgres", cnf.DataSourceName())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
