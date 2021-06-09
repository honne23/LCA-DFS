package directory

import "github.com/gofrs/uuid"

type Member interface {
	GetID() string
}
type Employee struct {
	id string
}

type Manager struct {
	Employee
	employees []Member
}

func (e *Employee) GetID() string {
	return e.id
}

func (e *Employee) ToManager(employees *[]Member) Manager {
	id := e.GetID()
	return NewManager(&id, employees)
}

func (m *Manager) GetEmployees() []Member {
	return m.employees
}
func (m *Manager) AddEmployees(employees []Member) {
	if len(m.employees) == 0 {
		m.employees = employees
	} else {
		m.employees = append(m.employees, employees...)
	}

}

//NewEmployee will automatically create an employee with a generated id
func NewEmployee() Employee {
	uuid, _ := uuid.NewV4()
	ID := uuid.String()
	return Employee{id: ID}
}

//NewManager will generate an ID if `id` is nil.
// A slice of employees can optionally be passed to this manager
func NewManager(id *string, employees *[]Member) Manager {
	var ID string
	if id != nil {
		ID = *id
	} else {
		uuid, _ := uuid.NewV4()
		ID = uuid.String()
	}
	var underlings []Member
	if employees != nil {
		underlings = *employees
	} else {
		underlings = make([]Member, 0)
	}
	return Manager{
		Employee: Employee{
			id: ID,
		},
		employees: underlings,
	}
}
