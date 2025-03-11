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

func (r *EmployeeRepository) SearchEmployeeByID(id string) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").First(&employee, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) SearchEmployeesByName(names string) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").
		Where("LOWER(employees.names) LIKE LOWER(?)", names+"%").
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

func (r *EmployeeRepository) UpdateUser(user *models.Employee) error {
	var existingUser models.User
	if err := r.DB.Preload("UserStateType").Preload("UserType").First(&existingUser, "id = ?", user.ID).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) CreateEmployee(employee models.Employee) (*models.Employee, error) {
	err := r.DB.Create(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
