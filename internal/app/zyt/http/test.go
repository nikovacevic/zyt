package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikovacevic/zyt/internal/app/zyt"
)

// TestController handles all testing routes
type TestController struct{}

// ApplyTestController creates a Test controller and applies the routes to the given Server
func ApplyTestController(server *Server) {
	NewTestController().Route(server)
}

// NewTestController creates a new TestController
func NewTestController() *TestController {
	return &TestController{}
}

// Route applies routes to the given Router
func (tc *TestController) Route(server *Server) {
	server.HandleFunc("/api/status", tc.Status()).Methods("GET")
	server.HandleFunc("/api/teapot", tc.Teapot()).Methods("GET")
	server.HandleFunc("/api/parrot/{message}", tc.Parrot()).Methods("GET")
}

// Status returns the API status
func (tc *TestController) Status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		WriteJSON(w, &zyt.Response{
			Errors:  nil,
			Message: "All good",
		})
	}
}

// Teapot returns the 418 I'm a teapot HTTP status
func (tc *TestController) Teapot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		WriteJSON(w, &zyt.Response{
			Errors: []error{
				fmt.Errorf("I'm a teapot"),
			},
			Message: "The server refuses to brew coffee because it is a teapot.",
		})
	}
}

// Parrot repeats what you said
func (tc *TestController) Parrot() http.HandlerFunc {
	// This only happens once
	fmt.Printf("Parrot is booting up\n")
	// Then the HandlerFunc is returned to run indefinitely
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Quiet! Parrot is talking!\n")
		vars := mux.Vars(r)
		payload := vars["message"]
		WriteJSON(w, &zyt.Response{
			Errors:  nil,
			Message: "Parrot is talking",
			Payload: payload,
		})
	}
}
