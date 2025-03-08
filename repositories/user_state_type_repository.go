package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type UserStateTypeRepository struct {
	DB *gorm.DB
}

func NewUserStateTypeRepository(db *gorm.DB) *UserStateTypeRepository {
	return &UserStateTypeRepository{DB: db}
}

func (r *UserStateTypeRepository) GetUserStateTypeByID(id string) (*models.UserStateType, error) {
	var UserStateType models.UserStateType
	err := r.DB.First(&UserStateType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &UserStateType, nil
}

func (r *UserStateTypeRepository) GetAllUserStateTypes() ([]models.UserStateType, error) {
	var UserStateTypes []models.UserStateType
	err := r.DB.Find(&UserStateTypes).Error
	if err != nil {
		return nil, err
	}
	return UserStateTypes, nil
}
