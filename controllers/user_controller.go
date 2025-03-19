package controllers

import (
	"errors"
	"net/http"

	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Service *services.UserService
	Auth    *utilities.AuthorizationUtil
}

func NewUserController(service *services.UserService, auth *utilities.AuthorizationUtil) *UserController {
	return &UserController{Service: service, Auth: auth}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_USER_BY_ID

	if !uc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	user, err := uc.Service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userDTO := dtos.GetUserDTO{
		ID:          user.ID,
		Email:       user.Email,
		Password:    user.Password,
		Token:       user.Token,
		UserTypeID:  user.UserTypeID,
		UserStateID: user.UserStateTypeID,
	}

	c.JSON(http.StatusOK, userDTO)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_USERS

	if !uc.Auth.CheckPermission(c, permissionId) {
		return
	}

	users, err := uc.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	var usersDTO []dtos.GetUserDTO
	for _, user := range users {

		userDTO := dtos.GetUserDTO{
			ID:          user.ID,
			Email:       user.Email,
			Password:    user.Password,
			Token:       user.Token,
			UserTypeID:  user.UserTypeID,
			UserStateID: user.UserStateTypeID,
		}

		usersDTO = append(usersDTO, userDTO)
	}

	c.JSON(http.StatusOK, usersDTO)
}

func (uc *UserController) SearchUsersByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_USER_BY_ID

	if !uc.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("id")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	users, err := uc.Service.SearchUsersByID(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	var usersDTO []dtos.GetUserDTO
	for _, user := range users {

		userDTO := dtos.GetUserDTO{
			ID:          user.ID,
			Email:       user.Email,
			Password:    user.Password,
			Token:       user.Token,
			UserTypeID:  user.UserTypeID,
			UserStateID: user.UserStateTypeID,
		}

		usersDTO = append(usersDTO, userDTO)
	}
	c.JSON(http.StatusOK, usersDTO)
}

func (uc *UserController) SearchUsersByEmail(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_USERS_BY_EMAIL

	if !uc.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("email")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	users, err := uc.Service.SearchUsersByEmail(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	var usersDTO []dtos.GetUserDTO
	for _, user := range users {

		userDTO := dtos.GetUserDTO{
			ID:          user.ID,
			Email:       user.Email,
			Password:    user.Password,
			Token:       user.Token,
			UserTypeID:  user.UserTypeID,
			UserStateID: user.UserStateTypeID,
		}

		usersDTO = append(usersDTO, userDTO)
	}
	c.JSON(http.StatusOK, usersDTO)
}

func (uc *UserController) UpdateUserState(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_USER_STATE

	if !uc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	var request struct {
		UserState int `json:"user_state"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.Service.UpdateUserState(id, request.UserState)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userDTO := dtos.GetUserDTO{
		ID:          user.ID,
		Email:       user.Email,
		Password:    user.Password,
		Token:       user.Token,
		UserTypeID:  user.UserTypeID,
		UserStateID: user.UserStateTypeID,
	}

	c.JSON(http.StatusOK, userDTO)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_USER

	if !uc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	var dto dtos.UpdateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.Service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user.Email = dto.Email
	user.Password = dto.Password
	user.Token = dto.Token
	user.UserTypeID = dto.UserTypeID
	user.UserStateTypeID = dto.UserStateID

	err = uc.Service.UpdateUser(user)
	var dtoUser dtos.GetUserDTO

	dtoUser.ID = user.ID
	dtoUser.Email = user.Email
	dtoUser.Password = user.Password
	dtoUser.Token = user.Token
	dtoUser.UserTypeID = user.UserTypeID
	dtoUser.UserStateID = user.UserStateTypeID

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, dtoUser)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_USER

	if !uc.Auth.CheckPermission(c, permissionId) {
		return
	}

	var dto dtos.CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, _ := uc.Service.GetUserByEmail(dto.Email)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		return
	}

	newUser := models.User{
		Email:           dto.Email,
		Password:        dto.Password,
		Token:           dto.Token,
		UserTypeID:      dto.UserTypeID,
		UserStateTypeID: dto.UserStateID,
	}

	createdUser, err := uc.Service.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userDTO := dtos.GetUserDTO{
		ID:          createdUser.ID,
		Email:       createdUser.Email,
		Password:    createdUser.Password,
		Token:       createdUser.Token,
		UserTypeID:  createdUser.UserTypeID,
		UserStateID: createdUser.UserStateTypeID,
	}

	c.JSON(http.StatusCreated, userDTO)
}
