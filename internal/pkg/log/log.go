package log

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	INFO = iota
	DEBUG
	WARNING
	ERROR
)

var prefix = map[int]string{
	INFO:    "INFO",
	DEBUG:   "DEBUG",
	WARNING: "WARNING",
	ERROR:   "ERROR",
}

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

// LogRequests wraps an http.Handler with per-request logging to stdout
func (l *Logger) LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		l.Printf(
			"%-20s %-8s %-8s %s",
			time.Now().Format("2006-01-02 15:04:05"),
			time.Since(start).Truncate(100*time.Microsecond),
			r.Method,
			r.RequestURI,
		)
	})
}

// Log wraps an http.Handler with per-request logging to stdout
func (l *Logger) Log(priority int, format string, v ...interface{}) {
	l.Printf(
		"%-29s %-8s %s",
		time.Now().Format("2006-01-02 15:04:05"),
		prefix[priority],
		fmt.Sprintf(format, v...),
	)
}
