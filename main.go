package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks/add", addTasksHandler)
	http.HandleFunc("/tasks/get", getTasksHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
