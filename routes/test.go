package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikovacevic/zyt-api/models"
)

// Test handles all testing routes
type Test struct{}

// TestRoutes creates a Test controller and applies the routes to the given Server
func TestRoutes(srv *Server) {
	NewTest().Route(srv)
}

// NewTest creates a new Test controller
func NewTest() *Test {
	return &Test{}
}

// Route applies routes to the given Router
func (tc *Test) Route(srv *Server) {
	srv.router.HandleFunc("/api/status", tc.Status()).Methods("GET")
	srv.router.HandleFunc("/api/teapot", tc.Teapot()).Methods("GET")
	srv.router.HandleFunc("/api/parrot/{message}", tc.Parrot()).Methods("GET")
}

// Status returns the API status
func (tc *Test) Status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		WriteJSON(w, &models.Response{
			Errors:  nil,
			Message: "All good",
		})
	}
}

// Teapot returns the 418 I'm a teapot HTTP status
func (tc *Test) Teapot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		WriteJSON(w, &models.Response{
			Errors: []error{
				fmt.Errorf("I'm a teapot"),
			},
			Message: "The server refuses to brew coffee because it is a teapot.",
		})
	}
}

// Parrot repeats what you said
func (tc *Test) Parrot() http.HandlerFunc {
	// This only happens once
	fmt.Printf("Parrot is booting up\n")
	// Then the HandlerFunc is returned to run indefinitely
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Quiet! Parrot is talking!\n")
		vars := mux.Vars(r)
		payload := vars["message"]
		WriteJSON(w, &models.Response{
			Errors:  nil,
			Message: "Parrot is talking",
			Payload: payload,
		})
	}
}
