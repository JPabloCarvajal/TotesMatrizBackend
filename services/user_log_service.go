package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type UserLogService struct {
	Repo *repositories.UserLogRepository
}

func NewUserLogService(repo *repositories.UserLogRepository) *UserLogService {
	return &UserLogService{Repo: repo}
}

// Obtener logs de un usuario
func (s *UserLogService) GetUserLogs(userID string) ([]models.UserLog, error) {
	return s.Repo.GetUserLogs(userID)
}
