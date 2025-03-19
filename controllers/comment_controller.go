package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	Service *services.CommentService
	Auth    *utilities.AuthorizationUtil
}

func NewCommentController(service *services.CommentService, auth *utilities.AuthorizationUtil) *CommentController {
	return &CommentController{Service: service, Auth: auth}
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PERMISSION_BY_ID

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

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
	permissionId := config.PERMISSION_GET_ALL_COMMENTS

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

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

	permissionId := config.PERMISSION_SEARCH_COMMENTS_BY_EMAIL

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

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
	permissionId := config.PERMISSION_CREATE_COMMENT

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

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
	permissionId := config.PERMISSION_UPDATE_COMMENT

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

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
