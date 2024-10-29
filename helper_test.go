package gotask

import (
	"fmt"
	"testing"
	"time"
)

// helper function, create structured task list and return slice of task ids
func createStructuredFileTaskRepositoryAndTaskIDSlice(t testing.TB, upperBound int) (FileTaskRepository, []uint64) {
	// Create a FileTaskRepository, create a slice of uint64 to save taskID's
	tl := NewFileTaskRepository()
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

func assertFileTaskRepositoryEquals(t *testing.T, got, want FileTaskRepository) {
	t.Helper()
	if got.Size != want.Size {
		t.Errorf("Size Error. Want: %d, Got: %d", want.Size, got.Size)
	}
	for index, task := range got.Tasks {
		assertTaskEquals(t, task, want.Tasks[index])
	}
}
