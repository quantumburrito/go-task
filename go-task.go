package go_task

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
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
	creationTime := time.Now().UTC() // using utc to ensure monolythic component is recovered
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
	t.Size += 1
	t.Tasks = append(t.Tasks, newTask)
}

func (t *TaskList) ReadFromFile(file *os.File) error {

	// reset file poniter to begining of file
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}

	// read file content
	bytesValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("unable to Read File: %w", err)
	}

	// unmarshel data into slice of tasks structs
	var tasks []Task
	err = json.Unmarshal(bytesValue, &tasks)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	t.Tasks = tasks
	t.Size = len(tasks)

	return nil

}

func (t *TaskList) WriteToFile(file *os.File) error {
	// clear the file before writing
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}

	jsonData, err := json.MarshalIndent(t.Tasks, "", "    ") /// pretty Printed Json
	if err != nil {
		return fmt.Errorf("failed to marshael Tasks: %w", err)
	}

	// write data to file
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write to File: %w", err)
	}

	return nil
}
