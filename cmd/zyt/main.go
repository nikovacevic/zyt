package main

import (
	"os"

	"github.com/nikovacevic/zyt/internal/app/zyt/http"
	"github.com/nikovacevic/zyt/internal/app/zyt/postgres"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

func main() {
	logger := log.NewLogger(os.Stdout)

	cnf, err := postgres.GetConfig("../..")
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v\n", err.Error())
	}

	db, err := postgres.NewDB(*cnf)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v\n", err.Error())
	}

	server := http.NewServer()
	http.NewAuthController(db, logger).Route(server)
	http.NewEventController(db, logger).Route(server)
	http.NewUserController(db, logger).Route(server)
	http.ListenAndServe(":1234", logger.LogRequests(server))
}
