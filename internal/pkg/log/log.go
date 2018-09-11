package log

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// INFO is used for basic information
	INFO = iota
	// DEBUG is used for debugging purposes
	DEBUG
	// WARNING is used to report warnings
	WARNING
	// ERROR is used to log errors
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

// New creates a new Logger
func New(out io.Writer) *Logger {
	return &Logger{
		Logger: *log.New(out, "", 0),
	}
}

// Fatal is equivalent to Log, followed by a call to os.Exit(1)
func (l *Logger) Fatal(priority int, format string, v ...interface{}) {
	s := fmt.Sprintf(
		"%-29s %-8s %s",
		time.Now().Format("2006-01-02 15:04:05"),
		prefix[priority],
		fmt.Sprintf(format, v...),
	)
	l.Print(s)
	os.Exit(1)
}

// Log prints the given formatting string and spread parameters in the manner
// of Printf to Logger, l, prepended with timestamp and priority
func (l *Logger) Log(priority int, format string, v ...interface{}) {
	l.Printf(
		"%-29s %-8s %s",
		time.Now().Format("2006-01-02 15:04:05"),
		prefix[priority],
		fmt.Sprintf(format, v...),
	)
}

// Panic is equivalent to Log, followed by a panic
func (l *Logger) Panic(priority int, format string, v ...interface{}) {
	s := fmt.Sprintf(
		"%-29s %-8s %s",
		time.Now().Format("2006-01-02 15:04:05"),
		prefix[priority],
		fmt.Sprintf(format, v...),
	)
	l.Print(s)
	panic(s)
}

// LogRequests wraps an http.Handler with per-request logging
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
