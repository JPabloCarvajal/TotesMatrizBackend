package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type CommentService struct {
	Repo *repositories.CommentRepository
}

func NewCommentService(repo *repositories.CommentRepository) *CommentService {
	return &CommentService{Repo: repo}
}

func (s *CommentService) GetCommentByID(id int) (*models.Comment, error) {
	return s.Repo.GetCommentByID(id)
}

func (s *CommentService) SearchCommentsByEmail(email string) ([]models.Comment, error) {
	return s.Repo.SearchCommentsByEmail(email)
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
	return s.Repo.GetAllComments()
}

func (s *CommentService) UpdateComment(comment *models.Comment) error {
	return s.Repo.UpdateComment(comment)
}

func (s *CommentService) CreateComment(comment models.Comment) (*models.Comment, error) {
	return s.Repo.CreateComment(&comment)
}

func (s *CommentService) SearchCommentsByID(query string) ([]models.Comment, error) {
	return s.Repo.SearchCommentsByID(query)
}

func (s *CommentService) SearchCommentsByName(name string) ([]models.Comment, error) {
	return s.Repo.SearchCommentsByName(name)
}
