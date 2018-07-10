package route

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nikovacevic/zyt-api/log"
	"github.com/nikovacevic/zyt-api/model"
	"github.com/nikovacevic/zyt-api/store"
)

// UserController handles all user routes
type UserController struct {
	UserStore store.User
	logger    *log.Logger
}

// NewUserController creates a new User controller
func NewUserController(es store.User, logger *log.Logger) *UserController {
	return &UserController{
		UserStore: es,
		logger:    logger,
	}
}

// Route applies routes to the given Router
func (ec *UserController) Route(srv *Server) {
	// Get an User
	srv.router.Handle("/api/user/{id}", ec.ViewUser()).Methods("GET")
}

// ViewUser retrieves and shows an User
func (ec *UserController) ViewUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			ec.logger.Printf("ERROR: failed to parse ID\n")
			fmt.Println(err)
			WriteJSON(w, &model.Response{
				Errors:  []error{fmt.Errorf("User %v not found", id)},
				Message: "User not found",
				Payload: nil,
			})
			return
		}

		user, err := ec.UserStore.ViewUser(id)
		if err != nil {
			ec.logger.Println(err)
		}

		if user == nil {
			WriteJSON(w, &model.Response{
				Errors:  []error{fmt.Errorf("User %v not found", id)},
				Message: "User not found",
				Payload: nil,
			})
			return
		}

		WriteJSON(w, &model.Response{
			Message: fmt.Sprintf("User %v", id.String()),
			Payload: user,
		})
		return
	})
}
