package services

import (
	"totesbackend/repositories"
)

type AuthorizationService struct {
	Repo *repositories.AuthorizationRepository
}

func NewAuthorizationService(repo *repositories.AuthorizationRepository) *AuthorizationService {
	return &AuthorizationService{Repo: repo}
}

func (s *AuthorizationService) UserHasPermission(email string, permissionID int) (bool, error) {
	return s.Repo.UserHasPermission(email, permissionID)
}
