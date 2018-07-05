package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server TODO
type Server struct {
	router *mux.Router
	// TODO db
}

// NewServer creates a Server
func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

// ServerHTTP serves over HTTP using the Server's router
func (srv *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	srv.router.ServeHTTP(w, req)
}
