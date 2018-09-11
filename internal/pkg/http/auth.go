package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nikovacevic/zyt/internal/app/zyt"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

// Auth middleware authenticates a request. If authentication fails, send a 401
// with WWW-Authorization header. If it succeeds, pass request through.
func Auth(sessionService zyt.SessionService) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
			if len(authHeader) != 2 {
				w.Header().Set("WWW-Authenticate", `Token realm="Restricted"`)
				http.Error(w, "Not authorized", http.StatusUnauthorized)
				return
			}

			token := []byte(authHeader[1])
			valid, err := sessionService.VerifySession(token)
			if !valid || err != nil {
				w.Header().Set("WWW-Authenticate", `Token realm="Restricted"`)
				http.Error(w, "Not authorized", http.StatusUnauthorized)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}

// AuthController handles all user routes
type AuthController struct {
	AuthService    zyt.AuthService
	SessionService zyt.SessionService
	logger         *log.Logger
}

// NewAuthController creates a new User controller
func NewAuthController(as zyt.AuthService, ss zyt.SessionService, logger *log.Logger) *AuthController {
	return &AuthController{
		AuthService:    as,
		SessionService: ss,
		logger:         logger,
	}
}

// Route applies routes to the given Router
func (ac *AuthController) Route(server *Server) {
	// Authenticates a User
	server.Handle("/api/authenticate", ac.Authenticate()).Methods("POST")
}

// Authenticate authenticates a user's credentials and returns a new session token
func (ac *AuthController) Authenticate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := ac.AuthService.AuthenticateUser(email, password)
		if err != nil || user == nil {
			WriteJSON(w, http.StatusUnauthorized, &zyt.Response{
				Errors:  []error{fmt.Errorf("Failed to authenticate")},
				Message: "Failed to authenticate",
			})
			return
		}

		session, err := ac.SessionService.CreateSession(user)
		if err != nil {
			HTTP500(w, "Failed to create session")
			return
		}

		HTTP200(w, fmt.Sprintf("Success"), struct {
			User  interface{} `json:"user"`
			Token string      `json:"token"`
		}{
			user,
			session.Token.String(),
		})
	})
}
