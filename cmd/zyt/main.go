package main

import (
	"fmt"
	"os"

	"github.com/nikovacevic/zyt/internal/pkg/http"
	"github.com/nikovacevic/zyt/internal/pkg/log"
	"github.com/nikovacevic/zyt/internal/pkg/postgres"
)

func main() {
	logger := log.New(os.Stdout)

	// Configure session
	// TODO
	// TODO pass to AuthService

	// Configure and connect to Postgres
	pgConfig, err := postgres.GetConfig("config", "postgres")
	if err != nil {
		logger.Fatal(log.ERROR, fmt.Sprintf("Postgres config: %s", err.Error()))
	}
	pgdb, err := postgres.NewDB(*pgConfig)
	if err != nil {
		logger.Fatal(log.ERROR, err.Error())
	}

	// Connect services to server and start listening
	server := http.NewServer()
	// http.NewAuthController(pgdb, errorService, jwtService, logger).Route(server)
	http.NewEventController(pgdb, logger).Route(server)
	http.NewUserController(pgdb, logger).Route(server)

	http.ListenAndServe(":1234", logger.LogRequests(server))
}
