package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (t *Task) add(a, b int) int {
	time.Sleep(10 * time.Second)
	fmt.Println(a + b)
	t.Status = "done"
	return a + b
}

func AddTask(task *Task) {
	task.Uuid = uuid.New()
	task.Status = "pending"

	if task.Body.FuncName == "add" {
		go task.add(task.Body.Input[0], task.Body.Input[1])
	}

	dataStore = append(dataStore, task)
}

func GetTasks() []*Task {
	return dataStore
}
