package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	Service *services.CommentService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{Service: service}
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	comment, err := cc.Service.GetCommentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	commentDTO := dtos.GetCommentDTO{
		ID:             comment.ID,
		Name:           comment.Name,
		LastName:       comment.LastName,
		Email:          comment.Email,
		Phone:          comment.Phone,
		ResidenceState: comment.ResidenceState,
		ResidenceCity:  comment.ResidenceCity,
		Comment:        comment.Comment,
	}

	c.JSON(http.StatusOK, commentDTO)
}

func (cc *CommentController) GetAllComments(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	comments, err := cc.Service.GetAllComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:             comment.ID,
			Name:           comment.Name,
			LastName:       comment.LastName,
			Email:          comment.Email,
			Phone:          comment.Phone,
			ResidenceState: comment.ResidenceState,
			ResidenceCity:  comment.ResidenceCity,
			Comment:        comment.Comment,
		})
	}

	c.JSON(http.StatusOK, commentsDTO)
}

func (cc *CommentController) SearchCommentsByEmail(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	comments, err := cc.Service.SearchCommentsByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search comments"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:             comment.ID,
			Name:           comment.Name,
			LastName:       comment.LastName,
			Email:          comment.Email,
			Phone:          comment.Phone,
			ResidenceState: comment.ResidenceState,
			ResidenceCity:  comment.ResidenceCity,
			Comment:        comment.Comment,
		})
	}

	c.JSON(http.StatusOK, commentsDTO)
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	var dto dtos.CreateCommentDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{
		Name:           dto.Name,
		LastName:       dto.LastName,
		Email:          dto.Email,
		Phone:          dto.Phone,
		ResidenceState: dto.ResidenceState,
		ResidenceCity:  dto.ResidenceCity,
		Comment:        dto.Comment,
	}

	createdComment, err := cc.Service.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	commentDTO := dtos.GetCommentDTO{
		ID:             createdComment.ID,
		Name:           createdComment.Name,
		LastName:       createdComment.LastName,
		Email:          createdComment.Email,
		Phone:          createdComment.Phone,
		ResidenceState: createdComment.ResidenceState,
		ResidenceCity:  createdComment.ResidenceCity,
		Comment:        createdComment.Comment,
	}

	c.JSON(http.StatusCreated, commentDTO)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var dto dtos.UpdateCommentDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.Service.GetCommentByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	comment.Name = dto.Name
	comment.LastName = dto.LastName
	comment.Email = dto.Email
	comment.Phone = dto.Phone
	comment.ResidenceState = dto.ResidenceState
	comment.ResidenceCity = dto.ResidenceCity
	comment.Comment = dto.Comment

	err = cc.Service.UpdateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	updatedCommentDTO := dtos.GetCommentDTO{
		ID:             comment.ID,
		Name:           comment.Name,
		LastName:       comment.LastName,
		Email:          comment.Email,
		Phone:          comment.Phone,
		ResidenceState: comment.ResidenceState,
		ResidenceCity:  comment.ResidenceCity,
		Comment:        comment.Comment,
	}

	c.JSON(http.StatusOK, updatedCommentDTO)
}

func (cc *CommentController) SearchCommentsByID(c *gin.Context) {
	query := c.Query("query")
	fmt.Println("Searching comments by ID with:", query)

	comments, err := cc.Service.SearchCommentsByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving comments"})
		return
	}

	if len(comments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No comments found"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:             comment.ID,
			Name:           comment.Name,
			LastName:       comment.LastName,
			Email:          comment.Email,
			Phone:          comment.Phone,
			ResidenceState: comment.ResidenceState,
			ResidenceCity:  comment.ResidenceCity,
			Comment:        comment.Comment,
		})
	}

	c.JSON(http.StatusOK, commentsDTO)
}

func (cc *CommentController) SearchCommentsByName(c *gin.Context) {
	query := c.Query("name")
	fmt.Println("Searching comments by name with:", query)

	comments, err := cc.Service.SearchCommentsByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving comments"})
		return
	}

	if len(comments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No comments found"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:             comment.ID,
			Name:           comment.Name,
			LastName:       comment.LastName,
			Email:          comment.Email,
			Phone:          comment.Phone,
			ResidenceState: comment.ResidenceState,
			ResidenceCity:  comment.ResidenceCity,
			Comment:        comment.Comment,
		})
	}

	c.JSON(http.StatusOK, commentsDTO)
}
