package vo

import (
	"time"
)

type TodoID string

func NewTodoID(randomStringID string, now time.Time, title, description string) (TodoID, error) {
	var obj = TodoID(randomStringID)
	return obj, nil
}

func (r TodoID) String() string {
	return string(r)
}
