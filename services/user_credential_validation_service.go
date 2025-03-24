package services

import (
	"totesbackend/repositories"
)

type UserCredentialValidationService struct {
	Repo *repositories.UserCredentialValidationRepository
}

func NewUserCredentialValidationService(repo *repositories.UserCredentialValidationRepository) *UserCredentialValidationService {
	return &UserCredentialValidationService{Repo: repo}
}

func (s *UserCredentialValidationService) ValidateUserCredentials(email, password string) (bool, error) {
	return s.Repo.ValidateUserCredentials(email, password)
}
