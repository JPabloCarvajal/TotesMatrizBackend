package services

import (
	"totesbackend/repositories"
)

type AuthorizationService struct {
	Repo     *repositories.AuthorizationRepository
	UserRepo *repositories.UserRepository
}

func NewAuthorizationService(repo *repositories.AuthorizationRepository, userRepo *repositories.UserRepository) *AuthorizationService {
	return &AuthorizationService{Repo: repo, UserRepo: userRepo}
}

func (s *AuthorizationService) UserHasPermission(email string, permissionID int) (bool, error) {
	return s.Repo.UserHasPermission(email, permissionID)
}
