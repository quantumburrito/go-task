package gotask

import (
	"math/rand"
	"time"
)

type Task struct {
	Description string
	Id          uint64
	Status      Status
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

func NewTask() Task {
	randNumber := rand.Uint64()
	creationTime := time.Now().UTC() // using utc to ensure monolythic component is recovered
	return Task{Description: "", Id: randNumber, Status: ToDo, CreatedAt: creationTime, ModifiedAt: creationTime}
}
