package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) GetCommentByID(id int) (*models.Comment, error) {
	var comment models.Comment
	err := r.DB.First(&comment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	err := r.DB.Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) SearchCommentsByEmail(email string) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.DB.Where("LOWER(email) LIKE LOWER(?)", email+"%").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	if err := r.DB.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepository) UpdateComment(comment *models.Comment) error {
	var existingComment models.Comment
	if err := r.DB.First(&existingComment, "id = ?", comment.ID).Error; err != nil {
		return err
	}

	return r.DB.Save(comment).Error
}
