package main

import (
	"net/http"

	"github.com/nikovacevic/zyt-api/routes"
)

func main() {
	srv := routes.NewServer()
	routes.TestRoutes(srv)
	http.ListenAndServe(":1234", srv)
}
