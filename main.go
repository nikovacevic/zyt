package main

import (
	"net/http"
	"os"

	"github.com/nikovacevic/zyt-api/log"
	"github.com/nikovacevic/zyt-api/route"
	"github.com/nikovacevic/zyt-api/store"
)

var db *store.DB
var logger *log.Logger

func init() {
	var err error

	logger = log.NewLogger(os.Stdout)

	db, err = store.NewDB("postgres", "postgres://niko@localhost/zyt?sslmode=disable")
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v\n", err.Error())
	}
}

func main() {
	srv := route.NewServer()
	route.NewEventController(db, logger).Route(srv)
	route.NewUserController(db, logger).Route(srv)
	route.NewTestController().Route(srv)
	http.ListenAndServe(":1234", logger.Request(srv))
}
