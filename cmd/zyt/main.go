package main

import (
	"os"

	"github.com/nikovacevic/zyt/internal/pkg/err"
	"github.com/nikovacevic/zyt/internal/pkg/http"
	"github.com/nikovacevic/zyt/internal/pkg/jwt"
	"github.com/nikovacevic/zyt/internal/pkg/log"
	"github.com/nikovacevic/zyt/internal/pkg/postgres"
)

func main() {
	logger := log.NewLogger(os.Stdout)
	errorService := err.NewService(logger)

	// Configure and connect to Postgres
	pgConfig, err := postgres.GetConfig("../..")
	errorService.CheckAndFatal(err)
	db, err := postgres.NewDB(*pgConfig)
	errorService.CheckAndFatal(err)

	// Configure JWT
	jwtConfig, err := jwt.GetConfig("../..")
	errorService.CheckAndFatal(err)
	jwtService := jwt.NewService(*jwtConfig)

	// Connect services to server and start listening
	server := http.NewServer()
	http.NewAuthController(db, errorService, jwtService, logger).Route(server)
	http.NewEventController(db, logger).Route(server)
	http.NewUserController(db, logger).Route(server)
	http.ListenAndServe(":1234", logger.LogRequests(server))
}
