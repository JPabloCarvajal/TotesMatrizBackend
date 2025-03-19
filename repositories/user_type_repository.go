package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type UserTypeRepository struct {
	DB *gorm.DB
}

func NewUserTypeRepository(db *gorm.DB) *UserTypeRepository {
	return &UserTypeRepository{DB: db}
}

func (r *UserTypeRepository) ObtainAllUserTypes() ([]models.UserType, error) {
	var userTypes []models.UserType
	err := r.DB.Preload("Roles").Find(&userTypes).Error
	if err != nil {
		return nil, err
	}
	return userTypes, nil
}

func (r *UserTypeRepository) GetUserTypeByID(id uint) (*models.UserType, error) {
	var userType models.UserType
	err := r.DB.Preload("Roles").First(&userType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &userType, nil
}

func (r *UserTypeRepository) Exists(userTypeID uint) (bool, error) {
	var count int64
	err := r.DB.Table("user_types").Where("id = ?", userTypeID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserTypeRepository) GetRolesForUserType(userTypeID uint) ([]uint, error) {
	var roleIDs []uint
	err := r.DB.Table("user_type_has_role").Select("role_id").
		Where("user_type_id = ?", userTypeID).
		Pluck("role_id", &roleIDs).Error
	if err != nil {
		return nil, err
	}
	return roleIDs, nil
}

func (r *UserTypeRepository) SearchUserTypesByID(query string) ([]models.UserType, error) {
	var userTypes []models.UserType
	err := r.DB.Preload("Roles").
		Where("CAST(id AS TEXT) LIKE ?", query+"%").
		Find(&userTypes).Error
	if err != nil {
		return nil, err
	}
	return userTypes, nil
}

func (r *UserTypeRepository) SearchUserTypesByName(query string) ([]models.UserType, error) {
	var userTypes []models.UserType
	err := r.DB.Preload("Roles").
		Where("LOWER(name) LIKE LOWER(?)", query+"%").
		Find(&userTypes).Error
	if err != nil {
		return nil, err
	}
	return userTypes, nil
}
