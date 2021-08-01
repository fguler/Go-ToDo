package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fguler/goToDo/domain"
	"github.com/gorilla/mux"
)

// GetTasks retuns all tasks
func (h *handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "tasks")

	tasks := h.taskService.GetTasks()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "something went bad", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

}

//FindByID returns a single task by ID
func (h *handler) FindByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	task, err := h.taskService.FindByID(vars["taskid"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "something went bad", http.StatusInternalServerError)
		log.Println(err)
		return
	}

}

//AddTask adds new task to the database
func (h *handler) AddTask(w http.ResponseWriter, r *http.Request) {

	// check if content-type is json
	if err := h.IfContentTypeIsJson(r.Header); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := domain.Task{}

	//using MaxBytesReader to enforce a maximum read of 1MB from response body
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// read task from body
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add the new task to database
	task, err := h.taskService.Add(task)
	if err != nil {
		http.Error(w, "something went bad", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// sent new task to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "something went bad", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// validate task
	/* 	if err := task.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} */

}

func (h *handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	//TODO
}
