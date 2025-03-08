package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	DB *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

func (r *PermissionRepository) GetPermissionByID(id uint) (*models.Permission, error) {
	var permission models.Permission
	err := r.DB.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepository) GetAllPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.DB.Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
