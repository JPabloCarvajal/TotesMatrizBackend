package services

import (
	"time"
	"totesbackend/models"
	"totesbackend/repositories"
)

type UserLogService struct {
	Repo *repositories.UserLogRepository
}

func NewUserLogService(repo *repositories.UserLogRepository) *UserLogService {
	return &UserLogService{Repo: repo}
}

func (s *UserLogService) CreateUserLog(userEmail, logMessage string) (*models.UserLog, error) {
	userLog := &models.UserLog{
		UserEmail: userEmail,
		Log:       logMessage,
		DateTime:  time.Now(),
	}

	return s.Repo.CreateUserLog(userLog)
}
