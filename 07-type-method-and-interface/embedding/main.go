package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return m.Reports
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "tâm",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	// var eFail Employee = m // compilation error!
	var eOK Employee = m.Employee // ok!
	fmt.Println(eOK.Description())
}
