package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) GetEmployeeByID(id string) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").First(&employee, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) SearchEmployeesByID(query string) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").
		Where("CAST(id AS TEXT) LIKE ?", query+"%").
		Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) SearchEmployeesByName(names string) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").
		Where("LOWER(names) LIKE LOWER(?)", names+"%").
		Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	var existingEmployee models.Employee
	if err := r.DB.Preload("User").Preload("IdentifierType").First(&existingEmployee, "id = ?", employee.ID).Error; err != nil {
		return err
	}

	if err := r.DB.Save(employee).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	if err := r.DB.Create(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}
