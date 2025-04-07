package services

import (
	"errors"
	"totesbackend/repositories"
	"totesbackend/services/utils"
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

func (s *AuthorizationService) LoginUser(email string, password string) error {

	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("invalid email or password")
	}

	if user.UserStateType.Name != "Active" {
		return errors.New("user is not active")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return errors.New("invalid email or password")
	}

	return nil
}
