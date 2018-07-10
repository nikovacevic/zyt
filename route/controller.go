package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Controller defines a collection of routes and handlers
type Controller interface {
	Route(*Server)
}

// WriteJSON writes JSON to the given ResponseWriter
func WriteJSON(w http.ResponseWriter, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
