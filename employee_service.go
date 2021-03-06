package services

import (
	"github.com/subodhqss/go-training/models"
	"github.com/subodhqss/go-training/repository"
)

type EmployeeService interface {
	PrintEmp() *models.Employee
}

type empSrv struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &empSrv{employeeRepo: employeeRepo}
}

func (emp *empSrv) PrintEmp() *models.Employee {
	employee := emp.employeeRepo.PrintEmployee()
	// test := &models.Test{ID: 1, Message: "Hi SUbodh"}
	return employee
}
