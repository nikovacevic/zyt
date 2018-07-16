package http

import (
	"fmt"
	"net/http"

	"github.com/nikovacevic/zyt/internal/app/zyt"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

// AuthController handles all user routes
type AuthController struct {
	AuthService  zyt.AuthService
	ErrorService zyt.ErrorService
	TokenService zyt.TokenService
	logger       *log.Logger
}

// NewAuthController creates a new User controller
func NewAuthController(as zyt.AuthService, es zyt.ErrorService, ts zyt.TokenService, logger *log.Logger) *AuthController {
	return &AuthController{
		AuthService:  as,
		ErrorService: es,
		TokenService: ts,
		logger:       logger,
	}
}

// Route applies routes to the given Router
func (ac *AuthController) Route(server *Server) {
	// Authenticates a User
	server.Handle("/api/authenticate", ac.Authenticate()).Methods("POST")
}

// Authenticate authenticates a user's credentials and returns and access token
func (ac *AuthController) Authenticate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := ac.AuthService.AuthenticateUser(email, password)
		ac.ErrorService.CheckAndLog(err)
		if user == nil {
			WriteJSON(w, &zyt.Response{
				Errors:  []error{fmt.Errorf("Incorrect username and password combination")},
				Message: "Incorrect username and password combination",
				Payload: nil,
			})
			return
		}

		token, err := ac.TokenService.GenerateToken(user)
		ac.ErrorService.CheckAndLog(err)
		if token == "" {
			WriteJSON(w, &zyt.Response{
				Errors:  []error{fmt.Errorf("Cannot issue access token")},
				Message: "Cannot issue access token",
				Payload: nil,
			})
			return
		}

		WriteJSON(w, &zyt.Response{
			Message: fmt.Sprintf("Authentication succeeded"),
			Payload: struct {
				User interface{} `json:"user"`
				JWT  string      `json:"jwt"`
			}{
				user,
				token,
			},
		})
		return
	})
}
