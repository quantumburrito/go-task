package go_task

import (
	"math/rand"
	"time"
)

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
