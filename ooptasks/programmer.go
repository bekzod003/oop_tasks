package ooptasks

import "errors"

// Database of tasks
// var tasks []Task

type Programmer struct {
	Id     int
	Name   string
	Status string
	Tasks  []*Task
}

var programmerIdCounter int

// Constructor
func NewProgrammer(name, status string) *Programmer {
	programmerIdCounter++

	return &Programmer{
		programmerIdCounter, name, status, make([]*Task, 0),
	}
}

// If programmer has more than 3 tasks at the moment
// this function return false, becaude one Programmer
// cant do more than 3 tasks
func (p *Programmer) isBisy() bool {
	doingCounter := 0
	for _, val := range p.Tasks {
		if val.Status == "doing" {
			doingCounter++
		}
	}

	if doingCounter > 3 {
		p.Status = "busy"
	} else {
		p.Status = "ready to take tasks"
	}
	return !(doingCounter > 3)
}

func (p *Programmer) Testing(taskId int) error {
	for i, task := range p.Tasks {
		if taskId == task.Id {
			p.Tasks[i].Status = "testing"
			return nil
		}
	}
	return errors.New("task witsh such id does not exist")
}

func (p *Programmer) Done(taskId int) error {
	for i, task := range p.Tasks {
		if taskId == task.Id {
			p.Tasks[i].Status = "done"
			return nil
		}
	}
	return errors.New("task witsh such id does not exist")
}
