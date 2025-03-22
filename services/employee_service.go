package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type EmployeeService struct {
	Repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.Repo.GetEmployeeByID(id)
}

func (s *EmployeeService) SearchEmployeesByID(query string) ([]models.Employee, error) {
	return s.Repo.SearchEmployeesByID(query)
}

func (s *EmployeeService) SearchEmployeesByName(names string) ([]models.Employee, error) {
	return s.Repo.SearchEmployeesByName(names)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAllEmployees()
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) error {
	return s.Repo.UpdateEmployee(employee)
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	return s.Repo.CreateEmployee(employee)
}
