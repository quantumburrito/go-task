package gotask

type Status string

const (
	ToDo       Status = "ToDo"
	InProgress Status = "In Progress"
	Done       Status = "Done"
)

func (s Status) IsValid() bool {
	switch s {
	case ToDo, InProgress, Done:
		return true
	default:
		return false
	}
}
