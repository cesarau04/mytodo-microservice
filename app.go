package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/melvinmt/firebase"
)

type task struct {
	id		int			`json:"id"`
	desc	string	`json:"desc"`
}

func getNoElementsFirebase() int {
	url := "https://mytodo-microservice.firebaseio.com/v1/indexValue"
	ref := firebase.NewReference(url).Export(false)
	var noElements int;
	if err := ref.Value(&noElements); err != nil {
		panic(err)
	}
	return noElements
}

func getAllTasks (w http.ResponseWriter, r *http.Request) {
	noElements := getNoElementsFirebase()
	url := "https://mytodo-microservice.firebaseio.com/v1/tasks/"

	var ph string;
	for i := 0; i < noElements; i++ {
		urlTemp := url + strconv.Itoa(i)
		ref := firebase.NewReference(urlTemp).Export(false)
		fmt.Println(ref.url)
		if err := ref.Value(&ph); err != nil {
			panic(err)
		}
		fmt.Fprintln(w,ph)
	}
}
func newTask (w http.ResponseWriter, r *http.Request) {
	nextNoElements := getNoElementsFirebase()
	url := "https://mytodo-microservice.firebaseio.com/v1/tasks/" + strconv.Itoa(nextNoElements)
	refTask := firebase.NewReference(url)
	if err := refTask.Write("some object here"); err != nil {
		panic(err)
	}
	url = "https://mytodo-microservice.firebaseio.com/v1/indexValue"
	refIndex := firebase.NewReference(url)
	if err := refIndex.Write(nextNoElements+1); err != nil {
		refTask.Delete()
		panic(err)
	}
}
func deleteTask (w http.ResponseWriter, r *http.Request) {
	url := "https://mytodo-microservice.firebaseio.com/v1/tasks/2" // + some variable passed
	ref := firebase.NewReference(url)
	if err := ref.Delete(); err != nil {
		panic(err)
	}
}
func getIdTask (w http.ResponseWriter, r *http.Request) {
	fmt.Println("TBI")
}

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks", newTask).Methods("POST")
	router.HandleFunc("/tasks", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", getIdTask).Methods("GET")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
