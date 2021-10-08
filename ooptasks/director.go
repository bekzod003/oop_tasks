package ooptasks

type Director struct {
	Id   int
	Name string
}

var directorIdCounter int

// Constructor
func NewDirector(name string) *Director {
	directorIdCounter++

	return &Director{
		directorIdCounter, name,
	}
}

// This function gives task to TeamLead and
// writes tasks to "Tasks"
func (d *Director) GiveTeamLeadTask(t *Task, tl *TeamLead) error {
	tl.Tasks = append(tl.Tasks, t)
	return nil
}
