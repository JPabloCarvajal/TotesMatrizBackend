package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type UserStateTypeService struct {
	Repo *repositories.UserStateTypeRepository
}

func NewUserStateTypeService(repo *repositories.UserStateTypeRepository) *UserStateTypeService {
	return &UserStateTypeService{Repo: repo}
}

func (s *UserStateTypeService) GetAllUserStateTypes() ([]models.UserStateType, error) {
	return s.Repo.GetAllUserStateTypes()
}

func (s *UserStateTypeService) GetUserStateTypeByID(id string) (*models.UserStateType, error) {
	return s.Repo.GetUserStateTypeByID(id)
}
