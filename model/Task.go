package model

import (
	"time"
)

/*
Task struct
*/
type Task struct {
	ID          int
	Name        string
	CreatedDate time.Time
	IsDone      bool
}
