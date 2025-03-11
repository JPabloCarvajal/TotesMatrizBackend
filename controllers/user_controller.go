package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
