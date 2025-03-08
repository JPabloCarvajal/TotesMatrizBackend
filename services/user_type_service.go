package services

import (
	"errors"
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
	userTypes, err := s.Repo.ObtainAllUserTypes()
	if err != nil {
		return nil, errors.New("error al obtener los tipos de usuario")
	}
	return userTypes, nil
}

func (s *UserTypeService) ObtainUserTypeByID(id uint) (*models.UserType, error) {
	userType, err := s.Repo.ObtainUserTypeByID(id)
	if err != nil {
		return nil, errors.New("tipo de usuario no encontrado")
	}
	return userType, nil
}

func (s *UserTypeService) GetRolesForUserType(userTypeID uint) ([]uint, error) {
	roleIDs, err := s.Repo.GetRolesForUserType(userTypeID)
	if err != nil {
		return nil, errors.New("error al obtener los roles del tipo de usuario")
	}
	return roleIDs, nil
}

func (s *UserTypeService) Exists(userTypeID uint) (bool, error) {
	exists, err := s.Repo.Exists(userTypeID)
	if err != nil {
		return false, errors.New("error al verificar la existencia del tipo de usuario")
	}
	return exists, nil
}
