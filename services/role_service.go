package services

import (
	"errors"
	"totesbackend/models"
	"totesbackend/repositories"
)

type RoleService struct {
	Repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	role, err := s.Repo.GetRoleByID(id)
	if err != nil {
		return nil, errors.New("rol no encontrado")
	}
	return role, nil
}

func (s *RoleService) GetRolePermissions(roleID uint) ([]uint, error) {
	permissions, err := s.Repo.GetRolePermissions(roleID)
	if err != nil {
		return nil, errors.New("error al obtener permisos del rol")
	}
	return permissions, nil
}

func (s *RoleService) GetAllPermissionsOfRole(roleID uint) ([]models.Permission, error) {
	permissions, err := s.Repo.GetAllPermissionsOfRole(roleID)
	if err != nil {
		return nil, errors.New("error al obtener la lista de permisos del rol")
	}
	return permissions, nil
}

func (s *RoleService) ExistRole(roleID uint) (bool, error) {
	exists, err := s.Repo.ExistRole(roleID)
	if err != nil {
		return false, errors.New("error al verificar la existencia del rol")
	}
	return exists, nil
}
