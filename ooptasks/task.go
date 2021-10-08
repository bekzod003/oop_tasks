package ooptasks

import (
	"fmt"
	"time"
)

type Task struct {
	Id        int
	Title     string
	GivenBy   TeamLead
	Deadline  string
	Status    string
	CreatedAt string
	UpdatedAt string
}

var taskIdCounter int

// Constructor
func NewTask(title string, gb TeamLead, dl string, st string) *Task {
	taskIdCounter++

	return &Task{
		taskIdCounter, title, gb, dl, st, fmt.Sprint(time.Now()), fmt.Sprint(time.Now()),
	}
}

func (t *Task) Update(st string) {
	t.Status = st
	t.UpdatedAt = fmt.Sprint(time.Now())
}
