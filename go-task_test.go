package go_task

import (
	"math/rand"
	"testing"
	"time"
)

func assertEquals(t *testing.T, field string, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("Field: %s, Wanted: %v, Got %v", field, got, want)
	}
}

func TestCreateTask(t *testing.T) {
	t.Run("Test Task Constructor", func(t *testing.T) {
		got := NewTask()
		currentTime := time.Now()
		want := Task{Description: "", Status: "ToDo", Id: rand.Uint64(), CreatedAt: currentTime, ModifiedAt: currentTime}

		if got.Description != want.Description {
			t.Errorf("Wanted: %s, Got: %s", want.Description, got.Description)
		}
		if got.Status != want.Status {
			t.Errorf("Wanted: %s, Got %s", want.Status, got.Status)
		}
		if got.Id != 0 {
			t.Errorf("Wanted: Int > 0, Got %d", got.Id)
		}
		if time.Second <= got.CreatedAt.Sub(want.CreatedAt).Abs() {
			t.Errorf("Wanted: %s, Got: %s", want.CreatedAt, got.CreatedAt)
		}
		if time.Second <= got.ModifiedAt.Sub(want.ModifiedAt).Abs() {
			t.Errorf("Wanted: %s, Got: %s", want.ModifiedAt, got.ModifiedAt)
		}

	})
}

func TestReadTodoFile(t *testing.T) {
	// TODO:: implement test to read mock todo file
}
