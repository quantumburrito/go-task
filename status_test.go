package gotask

import "testing"

func TestStatusType(t *testing.T) {
	t.Run("Test Default Status Creation", func(t *testing.T) {
		var status Status
		if status != "" {
			t.Errorf("Expected default Status to be empty, got %v", status)
		}
	})
	t.Run("Test Specific Status Values", func(t *testing.T) {
		status := ToDo
		if status != ToDo {
			t.Errorf("Expected status %v, but got %v", ToDo, status)
		}
		status = InProgress
		if status != InProgress {
			t.Errorf("Expected status %v, but got %v", InProgress, status)
		}

		status = Done
		if status != Done {
			t.Errorf("Expected status %v, but got %v", Done, status)
		}
	})
}
