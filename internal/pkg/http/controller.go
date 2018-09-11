package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nikovacevic/zyt/internal/app/zyt"

	"github.com/google/uuid"
)

// Controller defines a collection of routes and handlers
type Controller interface {
	Route(*http.Server)
}

// HTTP200 writes a "200 OK" response to the given response writer
// with a message and payload
func HTTP200(w http.ResponseWriter, message string, payload interface{}) {
	WriteJSON(w, http.StatusOK, &zyt.Response{
		Message: message,
		Payload: payload,
	})
}

// HTTP400 writes a "400 Bad Request" response to the given response writer
// with an appended message
func HTTP400(w http.ResponseWriter, message string) {
	s := fmt.Sprintf("Bad request: %s", message)
	WriteJSON(w, http.StatusBadRequest, &zyt.Response{
		Errors:  []error{errors.New(s)},
		Message: s,
		Payload: nil,
	})
}

// HTTP404 writes a "404 Not found" response to the given response writer
// with an appended message
func HTTP404(w http.ResponseWriter, message string) {
	s := fmt.Sprintf("Not found: %s", message)
	WriteJSON(w, http.StatusNotFound, &zyt.Response{
		Errors:  []error{errors.New(s)},
		Message: s,
		Payload: nil,
	})
}

// HTTP500 writes a "500 Internal Server Error" response to the given response
// writer with an error message
func HTTP500(w http.ResponseWriter, message string) {
	s := fmt.Sprintf("Internal Server Error: %s", message)
	WriteJSON(w, http.StatusInternalServerError, &zyt.Response{
		Errors:  []error{errors.New(s)},
		Message: s,
	})
}

// ParseFloat retrieves the request parameter with the given name and attempts
// to parse it
func ParseFloat(name string, r *http.Request) (float64, error) {
	vars := mux.Vars(r)
	s, ok := vars[name]
	if !ok {
		return 0, fmt.Errorf("Parameter %s not defined", name)
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// ParseInt ...
func ParseInt(name string, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	s, ok := vars[name]
	if !ok {
		return 0, fmt.Errorf("Parameter %s not defined", name)
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ParseString ...
func ParseString(name string, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	s, ok := vars[name]
	if !ok {
		return "", fmt.Errorf("Parameter %s not defined", name)
	}
	return s, nil
}

// ParseUUID ...
func ParseUUID(name string, r *http.Request) (*uuid.UUID, error) {
	vars := mux.Vars(r)
	s, ok := vars[name]
	if !ok {
		return nil, fmt.Errorf("Parameter %s not defined", name)
	}
	u, err := uuid.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("Invalid UUID format: %s", s)
	}
	return &u, nil
}

// WriteJSON writes JSON to the given ResponseWriter
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}
