package http

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikovacevic/zyt/internal/app/zyt"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

// UserController handles all user routes
type UserController struct {
	UserService zyt.UserService
	logger      *log.Logger
}

// NewUserController creates a new User controller
func NewUserController(us zyt.UserService, logger *log.Logger) *UserController {
	return &UserController{
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
		var id *uuid.UUID
		var user *zyt.User
		var err error

		if id, err = ParseUUID("id", r); err != nil {
			HTTP400(w, "Valid ID is required")
			return
		}

		if user, err = uc.UserService.ViewUser(*id); err != nil {
			uc.logger.Println(err)
			HTTP404(w, fmt.Sprintf("User %s", id))
			return
		}

		HTTP200(w, fmt.Sprintf("User %v", id.String()), struct {
			User *zyt.User `json:"user"`
		}{
			user,
		})
	})
}
