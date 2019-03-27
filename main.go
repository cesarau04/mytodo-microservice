package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Task : Desc-> Description of the task, Complete-> Whenever or not it is complete
type Task struct {
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
	Date  string `json:"date,omitempty"`
}

func main() {
	userPort := flag.Int("port", 9000, "port to use")
	flag.Parse()
	if *userPort < 1024 || *userPort > 65535 {
		log.Fatal("Port musb be > 1023 and < 65535")
	}
	log.Printf("PORT==> %d", *userPort)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks", newTask).Methods("POST")
	router.HandleFunc("/tasks", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks", markAsDone).Methods("PUT")
	router.HandleFunc("/tasks/{id}", getIDTask).Methods("GET")

	if err := http.ListenAndServe(":"+strconv.Itoa(*userPort), router); err != nil {
		log.Fatal(err)
	}
}
