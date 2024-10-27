package go_task

import (
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
	if got != want {
		t.Errorf("Wanted: %v, Got %v", want, got)
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
func TestReadTodoFile(t *testing.T) {
	// TODO:: implement test to read mock todo file
}
