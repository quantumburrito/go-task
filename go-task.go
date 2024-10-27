package go_task

import (
	"math/rand"
	"time"
)

const INITIAL_TASKLIST_CAPACITY = 100

type Task struct {
	Description string
	Id          uint64
	Status      string
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

func NewTask() Task {
	randNumber := rand.Uint64()
	creationTime := time.Now()
	return Task{Description: "", Id: randNumber, Status: "ToDo", CreatedAt: creationTime, ModifiedAt: creationTime}
}

type TaskList struct {
	Size  int
	Tasks []Task
}

func NewTaskList() TaskList {
	return TaskList{Size: 0, Tasks: make([]Task, 0, INITIAL_TASKLIST_CAPACITY)}
}

func (t *TaskList) AddTask(newTask Task) {
	t.Tasks = append(t.Tasks, newTask)
}
