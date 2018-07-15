package http

import (
	"fmt"
	"net/http"

	"github.com/nikovacevic/zyt/internal/app/zyt"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

// AuthController handles all user routes
type AuthController struct {
	AuthService zyt.AuthService
	logger      *log.Logger
}

// NewAuthController creates a new User controller
func NewAuthController(as zyt.AuthService, logger *log.Logger) *AuthController {
	return &AuthController{
		AuthService: as,
		logger:      logger,
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

		// TODO remove
		ac.logger.Log(log.INFO, "email: %s, password: %s", email, password)

		user, err := ac.AuthService.AuthenticateUser(email, password)
		if err != nil {
			ac.logger.Log(log.ERROR, err.Error())
		}
		if user == nil {
			WriteJSON(w, &zyt.Response{
				Errors:  []error{fmt.Errorf("Incorrect username and password combination")},
				Message: "Incorrect username and password combination",
				Payload: nil,
			})
			return
		}

		mockJWT := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwOi8vYXV0aGVudGljYXRpb24uc2VydmVyLmlvIiwic3ViIjoiMzQ3NTgyMzQ1MzQ1IiwibmFtZSI6IkJydWNlIFNjaG5laWVyIiwiZW1haWwiOiJicnVjZUBzY2huZWllci5pbyIsInRvdGFsX2Jvc3MiOnRydWUsImdlbml1c19sZXZlbCI6ImV4Y2VwdGlvbmFsIn0.0tp9kl09DqT53M1AxzvRFaKCZIa_nlLv9nvg-3uMvkU")
		WriteJSON(w, &zyt.Response{
			Message: fmt.Sprintf("Authentication succeeded"),
			Payload: struct {
				User interface{} `json:"user"`
				JWT  []byte      `json:"jwt"`
			}{
				user,
				mockJWT,
			},
		})
		return
	})
}
