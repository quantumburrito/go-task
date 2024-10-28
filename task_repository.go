package gotask

import "io"

const INITIAL_TASK_REPOSITORY_CAPACITY = 100

type TaskRepository interface {
	AddTask(newTask Task)                // Adds a new Task
	Retrieve(id uint64) (*Task, error)   // Finds and retrieves a task by ID
	RetrieveAll() ([]Task, error)        // Retrieves all tasks
	Update(task Task) error              // Updates an existing task
	Delete(id uint64) error              // Deletes a task by ID
	Load(source io.Reader) error         // Loads tasks from an external source (e.g., file, database)
	Persist(destination io.Writer) error // Persists tasks to an external destination
}
