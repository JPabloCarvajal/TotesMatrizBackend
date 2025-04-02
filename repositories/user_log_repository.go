package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type UserLogRepository struct {
	DB *gorm.DB
}

func NewUserLogRepository(db *gorm.DB) *UserLogRepository {
	return &UserLogRepository{DB: db}
}

func (r *UserLogRepository) CreateUserLog(userLog *models.UserLog) (*models.UserLog, error) {
	if err := r.DB.Create(userLog).Error; err != nil {
		return nil, err
	}
	return userLog, nil
}
