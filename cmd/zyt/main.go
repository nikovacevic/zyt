package main

import (
	"os"

	"github.com/nikovacevic/zyt/internal/app/zyt/http"
	"github.com/nikovacevic/zyt/internal/app/zyt/postgres"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

func main() {
	logger := log.NewLogger(os.Stdout)

	db, err := postgres.NewDB("postgres://niko@localhost/zyt?sslmode=disable")
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v\n", err.Error())
	}

	server := http.NewServer()
	http.NewEventController(db, logger).Route(server)
	http.NewUserController(db, db, logger).Route(server)
	http.ListenAndServe(":1234", logger.Request(server))
}
