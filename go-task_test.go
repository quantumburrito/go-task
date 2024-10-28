package go_task

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func assertEquals(t *testing.T, field string, got, want interface{}) {
	if got != want {
		t.Errorf("Field: %s, Wanted: %v, Got %v", field, got, want)
	}
}

func assertTimeCloseToNow(t *testing.T, name string, got time.Time) {
	if time.Second < time.Since(got) {
		t.Errorf("Expected %s to be close to now, got %s instead", name, got)
	}
}

func assertTaskEquals(t *testing.T, got, want Task) {
	assertEquals(t, "Created At", got.CreatedAt.String(), want.CreatedAt.String())
	assertEquals(t, "Modified At", got.ModifiedAt.String(), want.CreatedAt.String())
	assertEquals(t, "Status", got.Status, want.Status)
	assertEquals(t, "Description", got.Description, want.Description)
	assertEquals(t, "ID", got.Id, want.Id)
}

func assertTaskListEquals(t *testing.T, got, want TaskList) {
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
