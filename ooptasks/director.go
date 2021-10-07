package ooptasks

type Director struct {
	Id   int
	Name string
}

var directorIdCounter int

// Constructor
func NewDirector(id int, name string) *Director {
	directorIdCounter++

	return &Director{
		directorIdCounter, name,
	}
}

// This function gives task to Programmer and
// writes tasks to "Tasks"
func (d Director) GiveDevTask(t Task, p *Programmer) error {
	p.Tasks = append(p.Tasks, t)

	return nil
}
