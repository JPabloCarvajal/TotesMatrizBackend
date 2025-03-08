package services

import (
	"errors"
	"totesbackend/models"
	"totesbackend/repositories"
)

type PermissionService struct {
	Repo *repositories.PermissionRepository
}

func NewPermissionService(repo *repositories.PermissionRepository) *PermissionService {
	return &PermissionService{Repo: repo}
}

func (s *PermissionService) GetPermissionByID(id uint) (*models.Permission, error) {
	permission, err := s.Repo.GetPermissionByID(id)
	if err != nil {
		return nil, errors.New("permiso no encontrado")
	}
	return permission, nil
}

func (s *PermissionService) GetAllPermissions() ([]models.Permission, error) {
	permissions, err := s.Repo.GetAllPermissions()
	if err != nil {
		return nil, errors.New("error al obtener los permisos")
	}
	return permissions, nil
}
