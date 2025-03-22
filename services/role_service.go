package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type RoleService struct {
	Repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.Repo.GetAllRoles()
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.Repo.GetRoleByID(id)

}

func (s *RoleService) GetRolePermissions(id uint) ([]uint, error) {
	return s.Repo.GetRolePermissions(id)
}

func (s *RoleService) GetAllPermissionsOfRole(id uint) ([]models.Permission, error) {
	return s.Repo.GetAllPermissionsOfRole(id)
}

func (s *RoleService) ExistRole(id uint) (bool, error) {
	return s.Repo.ExistRole(id)
}

func (s *RoleService) SearchRolesByID(id string) ([]models.Role, error) {
	return s.Repo.SearchRolesByID(id)
}

func (s *RoleService) SearchRolesByName(name string) ([]models.Role, error) {
	return s.Repo.SearchRolesByName(name)
}
