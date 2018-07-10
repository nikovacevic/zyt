package route

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nikovacevic/zyt-api/log"
	"github.com/nikovacevic/zyt-api/model"
	"github.com/nikovacevic/zyt-api/store"
)

// EventController handles all event routes
type EventController struct {
	EventStore store.Event
	logger     *log.Logger
}

// NewEventController creates a new Event controller
func NewEventController(es store.Event, logger *log.Logger) *EventController {
	return &EventController{
		EventStore: es,
		logger:     logger,
	}
}

// Route applies routes to the given Router
func (ec *EventController) Route(srv *Server) {
	// Get an Event
	srv.router.Handle("/api/event/{id}", ec.ViewEvent()).Methods("GET")
	// Start a new Event
	srv.router.Handle("/api/event", ec.StartEvent()).Methods("POST")
	// End an Event
	srv.router.Handle("/api/event/{id}/end", ec.EndEvent()).Methods("POST")
}

// ViewEvent retrieves and shows an Event
func (ec *EventController) ViewEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			ec.logger.Printf("ERROR: failed to parse ID\n")
			fmt.Println(err)
			WriteJSON(w, &model.Response{
				Errors:  []error{fmt.Errorf("Event %v not found", id)},
				Message: "Event not found",
				Payload: nil,
			})
			return
		}

		event, err := ec.EventStore.ViewEvent(id)
		if err != nil {
			ec.logger.Println(err)
		}

		if event == nil {
			WriteJSON(w, &model.Response{
				Errors:  []error{fmt.Errorf("Event %v not found", id)},
				Message: "Event not found",
				Payload: nil,
			})
			return
		}

		WriteJSON(w, &model.Response{
			Message: fmt.Sprintf("Event %v", id.String()),
			Payload: event,
		})
		return
	})
}

// StartEvent starts a new blank Event
func (ec *EventController) StartEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// payload := vars["message"]
		// WriteJSON(w, &model.Response{
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
		// WriteJSON(w, &model.Response{
		// 	Errors:  nil,
		// 	Message: "Parrot is talking",
		// 	Payload: payload,
		// })
	})
}
