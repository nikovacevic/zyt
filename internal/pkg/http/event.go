package http

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikovacevic/zyt/internal/app/zyt"
	"github.com/nikovacevic/zyt/internal/pkg/log"
)

// EventController handles all event routes
type EventController struct {
	EventService zyt.EventService
	logger       *log.Logger
}

// NewEventController creates a new Event controller
func NewEventController(es zyt.EventService, logger *log.Logger) *EventController {
	return &EventController{
		EventService: es,
		logger:       logger,
	}
}

// Route applies routes to the given Router
func (ec *EventController) Route(server *Server) {
	// Get an Event
	server.Handle("/api/event/{id}", ec.ViewEvent()).Methods("GET")
	// Start a new Event
	server.Handle("/api/event", ec.StartEvent()).Methods("POST")
	// End an Event
	server.Handle("/api/event/{id}/end", ec.EndEvent()).Methods("POST")
}

// ViewEvent retrieves and shows an Event
func (ec *EventController) ViewEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id *uuid.UUID
		var event *zyt.Event
		var err error

		if id, err = ParseUUID("id", r); err != nil {
			HTTP400(w, "Valid ID is required")
			return
		}

		if event, err = ec.EventService.ViewEvent(*id); err != nil {
			ec.logger.Println(err)
			HTTP400(w, fmt.Sprintf("Event %s", id))
			return
		}

		HTTP200(w, fmt.Sprintf("Event %v", id.String()), struct {
			Event *zyt.Event `json:"event"`
		}{
			event,
		})
	})
}

// StartEvent starts a new blank Event
func (ec *EventController) StartEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// payload := vars["message"]
		// WriteJSON(w, &zyt.Response{
		// 	Errors:  nil,
		// 	Message: "Parrot is talking",
		// 	Payload: payload,
		// })
	})
}

// EndEvent ends an Event
func (ec *EventController) EndEvent() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// payload := vars["message"]
		// WriteJSON(w, &zyt.Response{
		// 	Errors:  nil,
		// 	Message: "Parrot is talking",
		// 	Payload: payload,
		// })
	})
}
