package ooptasks

// Database of tasks
var tasks []Task

type Programmer struct {
	Id    int
	Name  string
	Tasks []Task
}

var programmerIdCounter int

// Constructor
func NewProgrammer(id int, name string) *Programmer {
	programmerIdCounter++

	return &Programmer{
		programmerIdCounter, name, make([]Task, 0),
	}
}
