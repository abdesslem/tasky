package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tasky/core"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, it's running")
}

func addTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task core.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		panic(err)
	}
	if task.Body == nil {
		http.Error(w, "missing field 'body' from JSON object", http.StatusBadRequest)
		return
	}
	core.AddTask(&task)
	fmt.Fprintf(w, "Added tasks to data store%s !", task.Name)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	dataStore := core.GetTasks()
	err := json.NewEncoder(w).Encode(dataStore)
	if err != nil {
		log.Fatal("Error while encoding", err)
	}
}
