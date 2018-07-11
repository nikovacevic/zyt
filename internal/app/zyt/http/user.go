package http

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nikovacevic/zyt/internal/app/zyt"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

// UserController handles all user routes
type UserController struct {
	AuthService zyt.AuthService
	UserService zyt.UserService
	logger      *log.Logger
}

// NewUserController creates a new User controller
func NewUserController(as zyt.AuthService, us zyt.UserService, logger *log.Logger) *UserController {
	return &UserController{
		AuthService: as,
		UserService: us,
		logger:      logger,
	}
}

// Route applies routes to the given Router
func (uc *UserController) Route(server *Server) {
	// Get an User
	server.Handle("/api/user/{id}", uc.ViewUser()).Methods("GET")
}

// ViewUser retrieves and shows an User
func (uc *UserController) ViewUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			// Given ID is not a valid UUID
			uc.logger.Printf("ERROR: failed to parse ID\n")
			fmt.Println(err)
			WriteJSON(w, &zyt.Response{
				Errors:  []error{fmt.Errorf("User %v not found", id)},
				Message: "User not found",
				Payload: nil,
			})
			return
		}

		user, err := uc.UserService.ViewUser(id)
		if err != nil {
			uc.logger.Println(err)
		}
		if user == nil {
			WriteJSON(w, &zyt.Response{
				Errors:  []error{fmt.Errorf("User not found")},
				Message: "User not found",
				Payload: nil,
			})
			return
		}
		WriteJSON(w, &zyt.Response{
			Message: fmt.Sprintf("User %v", id.String()),
			Payload: struct {
				User interface{} `json:"user"`
			}{
				user,
			},
		})
		return
	})
}

// Authenticate authenticates a user's credentials and returns and access token
func (uc *UserController) Authenticate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email := vars["email"]
		password := vars["password"]

		user, err := uc.AuthService.AuthenticateUser(email, password)
		if err != nil {
			uc.logger.Println(err)
		}
		if user == nil {
			WriteJSON(w, &zyt.Response{
				Errors:  []error{fmt.Errorf("User not found")},
				Message: "User not found",
				Payload: nil,
			})
			return
		}
		WriteJSON(w, &zyt.Response{
			Message: fmt.Sprintf("Authentication succeeded"),
			Payload: struct {
				User interface{} `json:"user"`
			}{
				user,
			},
		})
		return
	})
}
