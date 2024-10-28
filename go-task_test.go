package go_task

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

// helper function, create structured task list and return slice of task ids
func createStructuredTaskListAndTaskIDSlice(t testing.TB, upperBound int) (TaskList, []uint64) {
	// Create a taskList, create a slice of uint64 to save taskID's
	tl := NewTaskList()
	taskIds := make([]uint64, 0, 10)

	// create 10 Tasks , save id's to slice
	for i := 0; i < 10; i++ {
		newTask := NewTask()
		newTask.Description = fmt.Sprintf("Task: %d", i)
		taskIds = append(taskIds, newTask.Id)
		tl.AddTask(newTask)
	}
	return tl, taskIds
}

func assertEquals(t *testing.T, field string, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("Field: %s, Wanted: %v, Got %v", field, got, want)
	}
}

func assertTimeCloseToNow(t *testing.T, name string, got time.Time) {
	t.Helper()
	if time.Second < time.Since(got) {
		t.Errorf("Expected %s to be close to now, got %s instead", name, got)
	}
}

func assertTaskEquals(t *testing.T, got, want Task) {
	t.Helper()
	assertEquals(t, "Created At", got.CreatedAt.String(), want.CreatedAt.String())
	assertEquals(t, "Modified At", got.ModifiedAt.String(), want.CreatedAt.String())
	assertEquals(t, "Status", got.Status, want.Status)
	assertEquals(t, "Description", got.Description, want.Description)
	assertEquals(t, "ID", got.Id, want.Id)
}

func assertTaskListEquals(t *testing.T, got, want TaskList) {
	t.Helper()
	if got.Size != want.Size {
		t.Errorf("Size Error. Want: %d, Got: %d", want.Size, got.Size)
	}
	for index, task := range got.Tasks {
		assertTaskEquals(t, task, want.Tasks[index])
	}
}

func TestCreateTask(t *testing.T) {
	t.Run("Test Task Constructor", func(t *testing.T) {
		got := NewTask()
		tests := []struct {
			field string
			got   interface{}
			want  interface{}
		}{
			{"Description", got.Description, ""},
			{"Status", got.Status, "ToDo"},
		}
		for _, tt := range tests {
			assertEquals(t, tt.field, tt.got, tt.want)
		}

		assertTimeCloseToNow(t, "Created At", got.CreatedAt)
		assertTimeCloseToNow(t, "Modified At", got.ModifiedAt)

	})
}

func TestTaskList(t *testing.T) {
	t.Run("Test Task List Constructor", func(t *testing.T) {
		taskList := NewTaskList()
		// Want empty task list with length 0 and no elements
		if taskList.Size != 0 {
			t.Errorf("Expected Size: 0, Got: %d", taskList.Size)
		}
		if len(taskList.Tasks) > 0 {
			t.Errorf("Expected length of task list: 0, Got: %d", len(taskList.Tasks))
		}
	})

	t.Run("Test AddTask method", func(t *testing.T) {
		taskList := NewTaskList()
		want := NewTask()
		taskList.AddTask(want)

		assertTaskEquals(t, taskList.Tasks[0], want)
	})
	t.Run("Add Multiple empty tasks to task list using AddTask", func(t *testing.T) {
		taskList := NewTaskList()

		// Generate 5 random Tasks
		testTasks := make([]Task, 5)
		for index, _ := range testTasks {
			testTasks[index] = NewTask()
		}

		// add them to the taskList
		for _, task := range testTasks {
			taskList.AddTask(task)
		}

		// assert added tasks are equivilant to expected value
		for index, task := range testTasks {
			assertTaskEquals(t, taskList.Tasks[index], task)
		}
	})
}

func TestTaskListFileIO(t *testing.T) {
	// Create a temporary File
	file, err := os.CreateTemp("", "example.tasks")
	if err != nil {
		log.Fatalf("Temp File %s could not be created", file.Name())
	}
	// at the end of the test, defer to removing the file
	defer os.Remove(file.Name())

	// Create TaskList and polulate with 5 random tasks
	taskList := NewTaskList()
	for i := 0; i < 5; i++ {
		task := NewTask()
		task.Description = fmt.Sprintf("Task Number: %d", i)
		taskList.AddTask(task)
	}

	// Write taskList to the temporary File
	if err := taskList.WriteToFile(file); err != nil {
		t.Fatalf("Failed to Write to file: %v", err)
	}

	// Read back into a new TaskList
	gotTaskList := NewTaskList()
	if err := gotTaskList.ReadFromFile(file); err != nil {
		t.Fatalf("Failed to Read from file: %v", err)
	}

	// Verify Contents
	assertTaskListEquals(t, gotTaskList, taskList)

}

func TestFindTask(t *testing.T) {

	t.Run("Test FindTask Method", func(t *testing.T) {

		tl, taskIds := createStructuredTaskListAndTaskIDSlice(t, 10)

		//randomly seelect 1 id from the slice
		randomIndex := rand.Intn(len(taskIds))
		randomTaskId := taskIds[randomIndex]

		// find random Task
		newTask, err := tl.FindTask(randomTaskId)
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
		tl, _ := createStructuredTaskListAndTaskIDSlice(t, 10)

		// generate random task id
		randomTaskId := rand.Uint64()

		// find task with this random id
		failedTask, err := tl.FindTask(randomTaskId)

		// check to see if findTask Failed, if not report error
		if err == nil {
			t.Errorf("FindTask did not fail as expected, %v", err)
		}

		// assert failed task equals empty task
		if failedTask != nil {
			t.Errorf("Find Task did not fail as expected, failedTask != nil")
		}

	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("Correct implementaion of Update Task", func(t *testing.T) {
		// create task list and task id slice with 10 random tasks
		tl, tIds := createStructuredTaskListAndTaskIDSlice(t, 10)

		// create a new task to update
		updatedTask := NewTask()
		updatedTask.Description = "Updated Task"
		updatedTask.Status = "Done"

		// select random task Id from tIds list, assign to updatedTask
		randomIndex := rand.Intn(len(tIds))
		updatedTask.Id = tIds[randomIndex]

		// update task with id
		err := tl.UpdateTask(updatedTask)
		if err != nil {
			t.Errorf("updateTask failed to update task: %v", err)
		}

	})

	t.Run("Incorrect Implementation of Update Task", func(t *testing.T) {
		// same as before, ignore task id slice
		tl, _ := createStructuredTaskListAndTaskIDSlice(t, 10)

		// create a new task to update
		updatedTask := NewTask()
		updatedTask.Description = "Updated Task"
		updatedTask.Status = "Done"

		// update task with id
		err := tl.UpdateTask(updatedTask)

		// if error is not encountered, fail
		if err == nil {
			t.Errorf("UpdateTask Expected to fail but didn't: %v", err)
		}
	})
}
