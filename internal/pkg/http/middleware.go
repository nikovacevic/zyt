package http

import "net/http"

// Middleware modifies the behavior of an http Handler, sitting between the
// request and the response to provide the ability to transform or early-return.
type Middleware func(http.Handler) http.Handler

// Apply applies one or more Middleware functions to an http Handler
func Apply(handler http.Handler, middleware ...Middleware) http.Handler {
	for _, mw := range middleware {
		handler = mw(handler)
	}
	return handler
}
