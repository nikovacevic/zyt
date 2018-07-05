package models

import (
	"encoding/json"
)

// Response represents the API status
type Response struct {
	Errors  []error
	Message string
	Payload interface{}
}

// MarshalJSON marshals an Response instance into JSON
func (r *Response) MarshalJSON() ([]byte, error) {
	var errors []string
	for _, err := range r.Errors {
		errors = append(errors, err.Error())
	}

	return json.Marshal(struct {
		Errors  []string    `json:"errors"`
		Message string      `json:"message"`
		Payload interface{} `json:"payload"`
	}{
		errors,
		r.Message,
		r.Payload,
	})
}
