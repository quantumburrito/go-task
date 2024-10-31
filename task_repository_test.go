package gotask

import "testing"

func TestCreate_Task(t *testing.T) {
	t.Run("Test Task Constructor", func(t *testing.T) {
		got := NewTask()
		tests := []struct {
			field string
			got   interface{}
			want  interface{}
		}{
			{"Description", got.Description, ""},
			{"Status", got.Status, ToDo},
		}
		for _, tt := range tests {
			assertEquals(t, tt.field, tt.got, tt.want)
		}

		assertTimeCloseToNow(t, "Created At", got.CreatedAt)
		assertTimeCloseToNow(t, "Modified At", got.ModifiedAt)

	})
}
