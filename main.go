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
	router.HandleFunc("/tasks/{id}", getIDTask).Methods("GET")

	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatal(err)
	}
}

// ctx := context.Background()
// conf := &firebase.Config{
// 	DatabaseURL: "https://mytodo-microservice.firebaseio.com/",
// }

// opt := option.WithCredentialsFile("key.json")

// app, err := firebase.NewApp(ctx, conf, opt)
// if err != nil {
// 	log.Fatalln("Error initializing app:", err)
// }

// client, err := app.Database(ctx)
// if err != nil {
// 	log.Fatalln("Error initializing database client:", err)
// }

/* Way to save 1
task := map[string]*Task{
	"26-03-2019--17:43:00": {
		Desc: "A beaitful task",
		Opcb: "Lmao",
	},
}
ref := client.NewRef("/")
taskRef := ref.Child("v1/tasks/")
if err := taskRef.Set(ctx, task); err != nil {
	log.Fatal("Error setting value:", err)
}
*/

/* Way to save 2
ref := client.NewRef("/v1/todo")
if _, err := ref.Push(ctx, &Task{
	Desc: "A bad task",
	Date: time.Now().Format("2006-01-02 15:04"),
	Done: true,
}); err != nil {
	log.Fatal("error while setting value", err)
}
*/
// ref := client.NewRef("/v1")
// var data map[string]interface{}
// if err := ref.Get(ctx, &data); err != nil {
// 	log.Fatalln("Error reading from database:", err)
// }
// fmt.Println(data)
