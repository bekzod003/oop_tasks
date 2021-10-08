package ooptasks

import (
	"errors"
	"fmt"
)

type TeamLead struct {
	Id   int
	Name string

	// programmers of programmer
	Assistents []*Programmer

	Tasks []*Task
}

var teamLeadIdCounter int

// Constructor
func NewTeamLead(name string) *TeamLead {
	teamLeadIdCounter++
	return &TeamLead{
		teamLeadIdCounter, name, make([]*Programmer, 0), make([]*Task, 0),
	}
}

func (tl *TeamLead) Delegate(t *Task) error {
	for _, p := range tl.Assistents {
		if p.isBisy() {
			theFreeProgrammer, err := tl.GetProgrammer(p.Id)
			if err != nil {
				fmt.Println(err)
				return err
			}

			theFreeProgrammer.Tasks = append(theFreeProgrammer.Tasks, t)
			return nil
		}
	}
	return errors.New("no free programmers")
}

func (tl *TeamLead) GetProgrammer(id int) (*Programmer, error) {
	for _, p := range tl.Assistents {
		if p.Id == id {
			return p, nil
		}
	}
	return &Programmer{}, errors.New("programmer with such id does not exist")
}

func (tl *TeamLead) FinishTask(taskId int) error {
	for _, task := range tl.Tasks {
		if task.Id == taskId && task.Status == "done" {
			task.Status = "finished"

			return nil
		}
	}
	return errors.New("task with such id does not exist, or status of task is not `done`")
}
