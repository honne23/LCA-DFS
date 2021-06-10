package directory

import "github.com/gofrs/uuid"

// This type generally represents a member of the company
type Member interface {
	//GetID returns the ID for a given node
	GetID() string
}

// An Employee is a base, concrete member
type Employee struct {
	id string
}

// A manager extends Employee to also have Employees asigned to them
type Manager struct {
	Employee
	employees []Member
}

func (e *Employee) GetID() string {
	return e.id
}

//ToManager converts a leaf node into a possible subtree
func (e *Employee) ToManager(employees *[]Member) Manager {
	id := e.GetID()
	return NewManager(&id, employees)
}

//GetEmployees will return all children of this node
func (m *Manager) GetEmployees() []Member {
	return m.employees
}

//AddEmployees adds children to the to the parent node
func (m *Manager) AddEmployees(employees []Member) {
	if len(m.employees) == 0 {
		m.employees = employees
	} else {
		m.employees = append(m.employees, employees...)
	}

}

//NewEmployee will automatically create an employee with a generated id
// This function will create a leaf node within the organization
func NewEmployee() Employee {
	uuid, _ := uuid.NewV4()
	ID := uuid.String()
	return Employee{id: ID}
}

//NewManager will generate an ID if `id` is nil.
// A slice of employees can optionally be passed to this manager.
// This function creates the root of a possible subtree within the organization.
func NewManager(id *string, employees *[]Member) Manager {
	var ID string
	if id != nil {
		ID = *id
	} else {
		uuid, _ := uuid.NewV4()
		ID = uuid.String()
	}
	var children []Member
	if employees != nil {
		children = *employees
	} else {
		children = make([]Member, 0)
	}
	return Manager{
		Employee: Employee{
			id: ID,
		},
		employees: children,
	}
}
