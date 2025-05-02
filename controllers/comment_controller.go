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
	Log     *utilities.LogUtil
}

func NewCommentController(service *services.CommentService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *CommentController {
	return &CommentController{Service: service, Auth: auth, Log: log}
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	idParam := c.Param("id")

	if cc.Log.RegisterLog(c, "Attempting to retrieve Comment with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_COMMENT_BY_ID
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for GetCommentByID")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Invalid comment ID format: "+idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	comment, err := cc.Service.GetCommentByID(id)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving Comment with ID "+idParam+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving comment"})
		return
	}

	if comment == nil {
		_ = cc.Log.RegisterLog(c, "Comment with ID "+idParam+" not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Successfully retrieved Comment with ID: "+idParam)

	commentDTO := dtos.GetCommentDTO{
		ID:       comment.ID,
		Name:     comment.Name,
		LastName: comment.LastName,
		Email:    comment.Email,
		Phone:    comment.Phone,
		Comment:  comment.Comment,
	}

	c.JSON(http.StatusOK, commentDTO)
}

func (cc *CommentController) GetAllComments(c *gin.Context) {
	if err := cc.Log.RegisterLog(c, "Attempting to retrieve all comments"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ALL_COMMENTS
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for GetAllComments")
		return
	}

	comments, err := cc.Service.GetAllComments()
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving all comments: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:       comment.ID,
			Name:     comment.Name,
			LastName: comment.LastName,
			Email:    comment.Email,
			Phone:    comment.Phone,
			Comment:  comment.Comment,
		})
	}

	_ = cc.Log.RegisterLog(c, "Successfully retrieved all comments")

	c.JSON(http.StatusOK, commentsDTO)
}

func (cc *CommentController) SearchCommentsByEmail(c *gin.Context) {
	if err := cc.Log.RegisterLog(c, "Attempting to search comments by email"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_COMMENTS_BY_EMAIL
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for SearchCommentsByEmail")
		return
	}

	email := c.Query("email")
	if email == "" {
		_ = cc.Log.RegisterLog(c, "Missing 'email' query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	comments, err := cc.Service.SearchCommentsByEmail(email)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error searching comments by email '"+email+"': "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search comments"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:       comment.ID,
			Name:     comment.Name,
			LastName: comment.LastName,
			Email:    comment.Email,
			Phone:    comment.Phone,
			Comment:  comment.Comment,
		})
	}

	_ = cc.Log.RegisterLog(c, "Successfully searched comments by email: "+email)

	c.JSON(http.StatusOK, commentsDTO)
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	if err := cc.Log.RegisterLog(c, "Attempting to create a comment"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_CREATE_COMMENT
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for CreateComment")
		return
	}

	var dto dtos.CreateCommentDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = cc.Log.RegisterLog(c, "Invalid input for CreateComment: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{
		Name:     dto.Name,
		LastName: dto.LastName,
		Email:    dto.Email,
		Phone:    dto.Phone,
		Comment:  dto.Comment,
	}

	createdComment, err := cc.Service.CreateComment(comment)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error creating comment: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	commentDTO := dtos.GetCommentDTO{
		ID:       createdComment.ID,
		Name:     createdComment.Name,
		LastName: createdComment.LastName,
		Email:    createdComment.Email,
		Phone:    createdComment.Phone,
		Comment:  createdComment.Comment,
	}

	_ = cc.Log.RegisterLog(c, "Successfully created comment with ID: "+strconv.Itoa(createdComment.ID))
	c.JSON(http.StatusCreated, commentDTO)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_COMMENT

	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for UpdateComment")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Invalid comment ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var dto dtos.UpdateCommentDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = cc.Log.RegisterLog(c, "Failed to bind JSON in UpdateComment: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.Service.GetCommentByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = cc.Log.RegisterLog(c, "Comment with ID "+strconv.Itoa(id)+" not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}
		_ = cc.Log.RegisterLog(c, "Internal error retrieving comment with ID "+strconv.Itoa(id)+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	comment.Name = dto.Name
	comment.LastName = dto.LastName
	comment.Email = dto.Email
	comment.Phone = dto.Phone
	comment.Comment = dto.Comment

	err = cc.Service.UpdateComment(comment)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Failed to update comment with ID "+strconv.Itoa(id)+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Successfully updated comment with ID: "+strconv.Itoa(id))

	updatedCommentDTO := dtos.GetCommentDTO{
		ID:       comment.ID,
		Name:     comment.Name,
		LastName: comment.LastName,
		Email:    comment.Email,
		Phone:    comment.Phone,
		Comment:  comment.Comment,
	}

	c.JSON(http.StatusOK, updatedCommentDTO)
}

func (cc *CommentController) SearchCommentsByID(c *gin.Context) {
	query := c.Query("id")

	permissionId := config.PERMISSION_SEARCH_COMMENTS_BY_ID

	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for SearchCommentsByID")
		return
	}

	comments, err := cc.Service.SearchCommentsByID(query)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving comments with ID "+query+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving comments"})
		return
	}

	if len(comments) == 0 {
		_ = cc.Log.RegisterLog(c, "No comments found for ID "+query)
		c.JSON(http.StatusNotFound, gin.H{"message": "No comments found"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:       comment.ID,
			Name:     comment.Name,
			LastName: comment.LastName,
			Email:    comment.Email,
			Phone:    comment.Phone,
			Comment:  comment.Comment,
		})
	}

	_ = cc.Log.RegisterLog(c, "Successfully retrieved comments with ID: "+query)
	c.JSON(http.StatusOK, commentsDTO)
}

func (cc *CommentController) SearchCommentsByName(c *gin.Context) {
	query := c.Query("name")

	permissionId := config.PERMISSION_SEARCH_COMMENTS_BY_NAME

	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for SearchCommentsByName")
		return
	}

	comments, err := cc.Service.SearchCommentsByName(query)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving comments with name "+query+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving comments"})
		return
	}

	if len(comments) == 0 {
		_ = cc.Log.RegisterLog(c, "No comments found for name "+query)
		c.JSON(http.StatusNotFound, gin.H{"message": "No comments found"})
		return
	}

	var commentsDTO []dtos.GetCommentDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.GetCommentDTO{
			ID:       comment.ID,
			Name:     comment.Name,
			LastName: comment.LastName,
			Email:    comment.Email,
			Phone:    comment.Phone,
			Comment:  comment.Comment,
		})
	}

	_ = cc.Log.RegisterLog(c, "Successfully retrieved comments with name: "+query)
	c.JSON(http.StatusOK, commentsDTO)
}
