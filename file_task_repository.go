package gotask

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type FileTaskRepository struct {
	Size     int
	Tasks    []Task
	filename string
	filepath string
}

func NewFileTaskRepository() FileTaskRepository {
	return FileTaskRepository{Size: 0, Tasks: make([]Task, 0, INITIAL_TASK_REPOSITORY_CAPACITY)}
}

func (t *FileTaskRepository) AddTask(newTask Task) {
	t.Size += 1
	t.Tasks = append(t.Tasks, newTask)
}

func (t *FileTaskRepository) Load(file *os.File) error {

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

func (t *FileTaskRepository) Persist(file *os.File) error {
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

func (t *FileTaskRepository) Retrieve(unknownID uint64) (*Task, error) {
	for _, task := range t.Tasks {
		if task.Id == unknownID {
			return &task, nil
		}
	}

	return nil, fmt.Errorf("task with id: %d not found", unknownID)
}

func (t *FileTaskRepository) Update(newTask Task) error {
	// check to see if task exists
	foundTask, err := t.Retrieve(newTask.Id)
	if err != nil {
		return err
	}
	foundTask.CreatedAt = newTask.CreatedAt
	foundTask.Description = newTask.Description
	foundTask.Id = newTask.Id
	foundTask.Status = newTask.Status

	return nil

}


func (t *FileTaskRepository) DescribeTasks(status string) (string, error) {
	taskList := ""
	var err error

	if t.Size == 0 {
		err = errors.New("can't list tasks of tasklist with size = 0")
		return taskList, err
	}
	if status == "" {
		for _, item := range t.Tasks {
			taskList += fmt.Sprintf("%d\t%s\t%s\t%s\t%s\n", item.Id, item.Description, item.Status, item.CreatedAt.String(), item.ModifiedAt.String())
		}
	} else {
		for _, item := range t.Tasks {
			if item.Status == status {
				taskList += fmt.Sprintf("%d\t%s\t%s\t%s\t%s\n", item.Id, item.Description, item.Status, item.CreatedAt.String(), item.ModifiedAt.String())
			}
		}
	}

	return taskList, err

}

func (t *FileTaskRepository) Delete(id uint64) error {
	task, err := t.Retrieve(id)
	if err != nil {
		return fmt.Errorf("Delete opeartion failed, %s", err)
	}

	// find the index of the task to delete
	index := -1
	for i, item := range t.Tasks {
		if item.Id == task.Id {
			index = i
			break
		}
	}

	// If the task was not found return an error
	if index == -1 {
		return fmt.Errorf("task with id %d not found", id)
	}

	// Delete the task from the slice
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)

	// change size after incrementation
	t.Size -= 1

	return nil
}
