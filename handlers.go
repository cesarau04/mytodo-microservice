package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

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

// Useful paths
var rootPathDB = "/"
var todoPathDB = "/v1/todo/"
var donePathDB = "/v1/done/"

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	todoRef := client.NewRef(todoPathDB)
	doneRef := client.NewRef(donePathDB)
	var todoTasks map[string]Task
	var doneTasks map[string]Task

	if err := todoRef.Get(ctx, &todoTasks); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Task to complete:\n")
	for key, value := range todoTasks {
		fmt.Fprintf(w, "\tID: %s\n\tTitle: %s\n\tDesc: %s\n\tDate: %s\n\n", key, value.Title, value.Desc, value.Date)
	}

	if err := doneRef.Get(ctx, &doneTasks); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "\nTask already done:\n")
	for key, value := range doneTasks {
		fmt.Fprintf(w, "\tID: %s\n\tTitle: %s\n\tDesc: %s\n\tDate: %s\n\n", key, value.Title, value.Desc, value.Date)
	}
}

func newTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	ref := client.NewRef(todoPathDB)

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Fatal(err)
	}

	newRef, err := ref.Push(ctx, task)
	if err != nil {
		log.Fatal("error while setting value", err)
	}
	fmt.Fprintf(w, "Successfully added with key: %s\n", newRef.Key)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var deleteID string

	if err := json.NewDecoder(r.Body).Decode(&deleteID); err != nil {
		log.Fatal(err)
	}
	ref := client.NewRef(todoPathDB + deleteID)
	if err := ref.Delete(ctx); err != nil {
		log.Fatal(err)
	}
	ref = client.NewRef(donePathDB + deleteID)
	if err := ref.Delete(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Deleted %s\n", deleteID)
}

func markAsDone(w http.ResponseWriter, r *http.Request) {
	var updateID string
	var dataToMove Task

	if err := json.NewDecoder(r.Body).Decode(&updateID); err != nil {
		log.Fatal(err)
	}
	ref := client.NewRef(todoPathDB + updateID)
	if err := ref.Get(ctx, &dataToMove); err != nil {
		log.Fatal(err)
	}
	if err := ref.Delete(ctx); err != nil {
		log.Fatal(err)
	}
	ref = client.NewRef(donePathDB + updateID)
	if err := ref.Set(ctx, &dataToMove); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Task %s, marked as done\n", updateID)
}

func getIDTask(w http.ResponseWriter, r *http.Request) {
	var data Task
	params := mux.Vars(r)
	taskID := params["id"]

	ref := client.NewRef(todoPathDB + taskID)
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Data of TaskID %s:\nTitle: %s\nDesc: %s\nDate: %s\n", taskID, data.Title, data.Desc, data.Date)
}
