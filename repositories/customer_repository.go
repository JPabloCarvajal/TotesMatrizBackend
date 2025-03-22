package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	err := r.DB.First(&customer, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.DB.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) GetCustomerByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	err := r.DB.First(&customer, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	if err := r.DB.Create(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) UpdateCustomer(customer *models.Customer) error {
	if err := r.DB.Save(customer).Error; err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) SearchCustomersByID(id string) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.DB.Where("CAST(id AS TEXT) LIKE ?", id+"%").Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) SearchCustomersByName(name string) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.DB.Where("LOWER(customer_name) LIKE LOWER(?)", name+"%").Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}
