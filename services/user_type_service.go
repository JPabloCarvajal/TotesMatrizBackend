package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type UserTypeService struct {
	Repo *repositories.UserTypeRepository
}

func NewUserTypeService(repo *repositories.UserTypeRepository) *UserTypeService {
	return &UserTypeService{Repo: repo}
}

func (s *UserTypeService) ObtainAllUserTypes() ([]models.UserType, error) {
	return s.Repo.ObtainAllUserTypes()
}

func (s *UserTypeService) GetUserTypeByID(id uint) (*models.UserType, error) {
	return s.Repo.GetUserTypeByID(id)
}

func (s *UserTypeService) GetRolesForUserType(userTypeID uint) ([]uint, error) {
	return s.Repo.GetRolesForUserType(userTypeID)
}

func (s *UserTypeService) Exists(userTypeID uint) (bool, error) {
	return s.Repo.Exists(userTypeID)
}

func (s *UserTypeService) SearchUserTypesByID(query string) ([]models.UserType, error) {
	return s.Repo.SearchUserTypesByID(query)
}

func (s *UserTypeService) SearchUserTypesByName(query string) ([]models.UserType, error) {
	return s.Repo.SearchUserTypesByName(query)
}
