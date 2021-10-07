package ooptasks

type TeamLead struct {
	Id   int
	Name string
}

var teamLeadIdCounter int

// Constructor
func NewTeamLead(id int, name string) *TeamLead {
	teamLeadIdCounter++
	return &TeamLead{
		teamLeadIdCounter, name,
	}
}

func (tl TeamLead) Delegate(p Programmer)
