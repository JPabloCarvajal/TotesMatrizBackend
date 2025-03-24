package repositories

import (
	"errors"
	"totesbackend/models"

	"gorm.io/gorm"
)

type UserCredentialValidationRepository struct {
	DB *gorm.DB
}

func NewUserCredentialValidationRepository(db *gorm.DB) *UserCredentialValidationRepository {
	return &UserCredentialValidationRepository{DB: db}
}

func (r *UserCredentialValidationRepository) ValidateUserCredentials(email, password string) (bool, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	return true, nil
}
