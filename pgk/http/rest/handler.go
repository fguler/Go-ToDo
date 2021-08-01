package rest

import (
	"errors"
	"net/http"

	"github.com/fguler/goToDo/pgk/task"
	"github.com/gorilla/mux"
)

type handler struct {
	taskService *task.Service
}

func newHandler(ts *task.Service) *handler {
	return &handler{
		taskService: ts,
	}
}

func RegisterRoutes(ts *task.Service, r *mux.Router) {

	h := newHandler(ts)

	r.HandleFunc("/api/v1/tasks", h.GetTasks).Methods("GET")
	r.HandleFunc("/api/v1/task", h.AddTask).Methods("POST")
	r.HandleFunc("/api/v1/task/{taskid}", h.FindByID).Methods("GET")

}

func (h handler) IfContentTypeIsJson(rh http.Header) error {
	ct := rh.Get("Content-Type")
	if ct == "application/json" {
		return nil
	}
	return errors.New("expected Content-Type header is application/json")
}
