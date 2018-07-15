package main

import (
	"os"

	"github.com/nikovacevic/zyt/internal/pkg/err"
	"github.com/nikovacevic/zyt/internal/pkg/http"
	"github.com/nikovacevic/zyt/internal/pkg/log"
	"github.com/nikovacevic/zyt/internal/pkg/postgres"
)

func main() {
	logger := log.NewLogger(os.Stdout)

	errorService := err.NewService(logger)

	cnf, err := postgres.GetConfig("../..")
	errorService.CheckAndFatal(err)

	db, err := postgres.NewDB(*cnf)
	errorService.CheckAndFatal(err)

	server := http.NewServer()
	http.NewAuthController(db, logger).Route(server)
	http.NewEventController(db, logger).Route(server)
	http.NewUserController(db, logger).Route(server)
	http.ListenAndServe(":1234", logger.LogRequests(server))
}
