package services

import (
	"errors"
	"totesbackend/repositories"
	"totesbackend/services/utils"
)

type UserCredentialValidationService struct {
	UserRepo *repositories.UserRepository
}

func NewUserCredentialValidationService(userRepo *repositories.UserRepository) *UserCredentialValidationService {
	return &UserCredentialValidationService{UserRepo: userRepo}
}

func (s *UserCredentialValidationService) ValidateUserCredentials(email, password string) error {
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
