package log

import (
	"io"
	"log"
	"net/http"
	"time"
)

// Logger is a logger
type Logger struct {
	log.Logger
}

// NewLogger creates a new Logger
func NewLogger(out io.Writer) *Logger {
	return &Logger{
		Logger: *log.New(out, "", 0),
	}
}

// Request wraps a HandlerFunc with logging to stdout
func (l *Logger) Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		l.Printf(
			"%-6s\t%-30s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

// Error wraps an erroneous request's HandlerFunc with logging to stdout
func (l *Logger) Error(next http.Handler, code string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		l.Printf(
			"%-6s\t%-30s\t%s",
			code,
			r.RequestURI,
			time.Since(start),
		)
	})
}
