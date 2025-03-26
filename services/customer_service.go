package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type CustomerService struct {
	Repo *repositories.CustomerRepository
}

func NewCustomerService(repo *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{Repo: repo}
}

func (s *CustomerService) GetCustomerByID(id int) (*models.Customer, error) {
	return s.Repo.GetCustomerByID(id)
}

func (s *CustomerService) GetCustomerByCustomerID(customerID string) (*models.Customer, error) {
	return s.Repo.GetCustomerByCustomerID(customerID)
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	return s.Repo.GetAllCustomers()
}

func (s *CustomerService) GetCustomerByEmail(email string) (*models.Customer, error) {
	return s.Repo.GetCustomerByEmail(email)
}

func (s *CustomerService) CreateCustomer(customer models.Customer) (*models.Customer, error) {
	return s.Repo.CreateCustomer(&customer)
}

func (s *CustomerService) UpdateCustomer(customer *models.Customer) error {
	return s.Repo.UpdateCustomer(customer)
}

func (s *CustomerService) SearchCustomersByID(id string) ([]models.Customer, error) {
	return s.Repo.SearchCustomersByID(id)
}

func (s *CustomerService) SearchCustomersByName(name string) ([]models.Customer, error) {
	return s.Repo.SearchCustomersByName(name)
}

func (s *CustomerService) SearchCustomersByLastName(lastname string) ([]models.Customer, error) {
	return s.Repo.SearchCustomersByLastName(lastname)
}
