package repository

import (
	"github.com/subodhqss/go-training/models"
)

type EmployeeRepository interface {
	PrintEmployee() *models.Employee
}

func NewEmpRepo() EmployeeRepository {
	return &employeeRepo{}
}

type employeeRepo struct{}

func (tr *employeeRepo) PrintEmployee() *models.Employee {
	employee := &models.Employee{Name: "Rajnish", Add: "Noida", Designation: "Software Engg", Company_name: "Qss Org", Date_of_joining: "18/06/2021", Emp_code: 1002}

	return employee
}
