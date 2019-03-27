package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Task : Desc-> Description of the task, Complete-> Whenever or not it is complete
type Task struct {
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
	Date  string `json:"date,omitempty"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks", newTask).Methods("POST")
	router.HandleFunc("/tasks", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks", markAsDone).Methods("PUT")
	router.HandleFunc("/tasks/{id}", getIDTask).Methods("GET")

	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatal(err)
	}
}
