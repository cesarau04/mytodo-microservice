package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Firebase setup
var ctx = context.Background()
var conf = &firebase.Config{
	DatabaseURL: "https://mytodo-microservice.firebaseio.com/",
}
var opt = option.WithCredentialsFile("key.json")
var app, _ = firebase.NewApp(ctx, conf, opt)
var client, _ = app.Database(ctx)

var rootPathDB = "/"
var todoPathDB = "/v1/todo/"
var donePathDB = "/v1/done/"

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	// Here goes the logic to call Firebase to retrieve the tasks from DB.
	todoRef := client.NewRef(todoPathDB)
	doneRef := client.NewRef(donePathDB)
	var todoTasks map[string]Task
	var doneTasks map[string]Task

	if err := todoRef.Get(ctx, &todoTasks); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Task to complete:\n")
	for key, value := range todoTasks {
		fmt.Fprintf(w, "ID: %s\nTitle: %s\nDesc: %s\nDate: %s\n", key, value.Title, value.Desc, value.Date)
	}

	if err := doneRef.Get(ctx, &doneTasks); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Task already done\n")
	for key, value := range doneTasks {
		fmt.Fprintf(w, "ID: %s\nTitle: %s\nDesc: %s\nDate: %s\n", key, value.Title, value.Desc, value.Date)
	}
}

func newTask(w http.ResponseWriter, r *http.Request) {

}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete")
}
func getIDTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getByID")
}
