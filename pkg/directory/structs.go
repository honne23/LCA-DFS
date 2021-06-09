package directory

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
