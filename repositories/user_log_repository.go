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

func (r *UserLogRepository) CreateUserLog(log *models.UserLog) error {
	return r.DB.Create(log).Error
}

func (r *UserLogRepository) GetUserLogs(userID string) ([]models.UserLog, error) {
	var userLogs []models.UserLog
	err := r.DB.Where("user_id = ?", userID).Order("date_time DESC").Find(&userLogs).Error
	return userLogs, err
}
