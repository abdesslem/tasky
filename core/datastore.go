package core

import (
	"github.com/google/uuid"
)

type Task struct {
	Name   string    `json:"name"`
	Retry  bool      `json:"retry"`
	Body   *TaskBody `json:"body"`
	Status string    `json:"status"`
	Uuid   uuid.UUID `json:"uuid"`
}

type TaskBody struct {
	FuncName string `json:"func_name"`
	Input    []int  `json:"input"`
}

var dataStore []*Task
