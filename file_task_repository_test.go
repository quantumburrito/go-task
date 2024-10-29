package gotask

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
)

func TestFileTaskRepository(t *testing.T) {
	t.Run("Test Task List Constructor", func(t *testing.T) {
		FileTaskRepository := NewFileTaskRepository()
		// Want empty task list with length 0 and no elements
		if FileTaskRepository.Size != 0 {
			t.Errorf("Expected Size: 0, Got: %d", FileTaskRepository.Size)
		}
		if len(FileTaskRepository.Tasks) > 0 {
			t.Errorf("Expected length of task list: 0, Got: %d", len(FileTaskRepository.Tasks))
		}
	})

	t.Run("Test AddTask method", func(t *testing.T) {
		FileTaskRepository := NewFileTaskRepository()
		want := NewTask()
		FileTaskRepository.AddTask(want)

		assertTaskEquals(t, FileTaskRepository.Tasks[0], want)
	})
	t.Run("Add Multiple empty tasks to task list using AddTask", func(t *testing.T) {
		FileTaskRepository := NewFileTaskRepository()

		// Generate 5 random Tasks
		testTasks := make([]Task, 5)
		for index, _ := range testTasks {
			testTasks[index] = NewTask()
		}

		// add them to the FileTaskRepository
		for _, task := range testTasks {
			FileTaskRepository.AddTask(task)
		}

		// assert added tasks are equivilant to expected value
		for index, task := range testTasks {
			assertTaskEquals(t, FileTaskRepository.Tasks[index], task)
		}
	})
}

func TestFileTaskRepositoryFileIO(t *testing.T) {
	// Create a temporary File
	file, err := os.CreateTemp("", "example.tasks")
	if err != nil {
		log.Fatalf("Temp File %s could not be created", file.Name())
	}
	// at the end of the test, defer to removing the file
	defer os.Remove(file.Name())

	// Create FileTaskRepository and polulate with 5 random tasks
	FileTaskRepository := NewFileTaskRepository()
	for i := 0; i < 5; i++ {
		task := NewTask()
		task.Description = fmt.Sprintf("Task Number: %d", i)
		FileTaskRepository.AddTask(task)
	}

	// Write FileTaskRepository to the temporary File
	if err := FileTaskRepository.Persist(file); err != nil {
		t.Fatalf("Failed to Write to file: %v", err)
	}

	// Read back into a new FileTaskRepository
	gotFileTaskRepository := NewFileTaskRepository()
	if err := gotFileTaskRepository.Load(file); err != nil {
		t.Fatalf("Failed to Read from file: %v", err)
	}

	// Verify Contents
	assertFileTaskRepositoryEquals(t, gotFileTaskRepository, FileTaskRepository)

}

func TestRetrieve_FileTaskRepository(t *testing.T) {

	t.Run("Test Retrieve Method", func(t *testing.T) {

		tl, taskIds := createStructuredFileTaskRepositoryAndTaskIDSlice(t, 10)

		//randomly seelect 1 id from the slice
		randomIndex := rand.Intn(len(taskIds))
		randomTaskId := taskIds[randomIndex]

		// find random Task
		newTask, err := tl.Retrieve(randomTaskId)
		if err != nil {
			t.Errorf("couldn't find task: %v \t Error: %v", newTask, err)
		}

		// check that new Task Id is correct ID
		if newTask.Id != randomTaskId {
			t.Errorf("Wanted: %d, Got: %d", randomTaskId, newTask.Id)
		}

	})

	t.Run("Test that Find Task Method Fails when no task is found", func(t *testing.T) {

		// create structured task list, don't need task id slice
		tl, _ := createStructuredFileTaskRepositoryAndTaskIDSlice(t, 10)

		// generate random task id
		randomTaskId := rand.Uint64()

		// find task with this random id
		failedTask, err := tl.Retrieve(randomTaskId)

		// check to see if Retrieve Failed, if not report error
		if err == nil {
			t.Errorf("Retrieve did not fail as expected, %v", err)
		}

		// assert failed task equals empty task
		if failedTask != nil {
			t.Errorf("Find Task did not fail as expected, failedTask != nil")
		}

	})
}

func TestUpdate_FileTaskRepository(t *testing.T) {
	t.Run("Correct implementaion of Update Task", func(t *testing.T) {
		// create task list and task id slice with 10 random tasks
		tl, tIds := createStructuredFileTaskRepositoryAndTaskIDSlice(t, 10)

		// create a new task to update
		updatedTask := NewTask()
		updatedTask.Description = "Updated Task"
		updatedTask.Status = "Done"

		// select random task Id from tIds list, assign to updatedTask
		randomIndex := rand.Intn(len(tIds))
		updatedTask.Id = tIds[randomIndex]

		// update task with id
		err := tl.Update(updatedTask)
		if err != nil {
			t.Errorf("Update failed to update task: %v", err)
		}

	})

	t.Run("Incorrect Implementation of Update Task", func(t *testing.T) {
		// same as before, ignore task id slice
		tl, _ := createStructuredFileTaskRepositoryAndTaskIDSlice(t, 10)

		// create a new task to update
		updatedTask := NewTask()
		updatedTask.Description = "Updated Task"
		updatedTask.Status = "Done"

		// update task with id
		err := tl.Update(updatedTask)

		// if error is not encountered, fail
		if err == nil {
			t.Errorf("Update Expected to fail but didn't: %v", err)
		}
	})
}

// Test List All Tasks
func TestDescribeTasks_TaskRepository(t *testing.T) {
	t.Run("Print Tasks in newly created task list", func(t *testing.T) {
		tl, _ := createStructuredFileTaskRepositoryAndTaskIDSlice(t, 10)

		got, err := tl.DescribeTasks()
		if err != nil {
			t.Errorf("Couldn't List all Tasks: %v", err)
		}
		var want string

		for _, item := range tl.Tasks {
			want += fmt.Sprintf("%d\t%s\t%s\t%s\t%s\n", item.Id, item.Description, item.Status, item.CreatedAt.String(), item.ModifiedAt.String())
		}

		if want != got {
			t.Errorf("Want:\t%s \n Got:\t%s\n", want, got)
		}
	})
	t.Run("Attempt to list tasks with ", func(t *testing.T) {
		tl := NewFileTaskRepository()
		got, err := tl.DescribeTasks()
		if err == nil {
			t.Errorf("trying to list tasks from task repository with no tasks did not return err: %v", err)
		}
		if got != "" {
			t.Errorf("trying to list tasks from task repository with size=0 did not return empty string")
		}

	})

}
