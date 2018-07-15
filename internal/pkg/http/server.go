package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server serves HTTP
type Server struct {
	*mux.Router
}

// NewServer creates a Server
func NewServer() *Server {
	return &Server{
		mux.NewRouter(),
	}
}

// ListenAndServe proxies the net/http package ListenAndServe
func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
