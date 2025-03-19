package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Service *services.CustomerService
	Auth    *utilities.AuthorizationUtil
}

func NewCustomerController(service *services.CustomerService, auth *utilities.AuthorizationUtil) *CustomerController {
	return &CustomerController{Service: service, Auth: auth}
}

func (cc *CustomerController) GetAllCustomers(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_CUSTOMERS

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

	customers, err := cc.Service.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) GetCustomerByID(c *gin.Context) {

	permissionId := config.PERMISSION_GET_CUSTOMER_BY_ID

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	customer, err := cc.Service.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_CUSTOMER

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

	var dto dtos.CreateCustomerDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	customer := models.Customer{
		CustomerName:     dto.CustomerName,
		CustomerId:       dto.CustomerId,
		IsBusiness:       dto.IsBusiness,
		Address:          dto.Address,
		PhoneNumbers:     dto.PhoneNumbers,
		CustomerState:    dto.CustomerState,
		Email:            dto.Email,
		LastName:         dto.LastName,
		IdentifierTypeID: dto.IdentifierTypeID,
	}

	createdCustomer, err := cc.Service.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating customer"})
		return
	}

	c.JSON(http.StatusCreated, createdCustomer)
}

func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_CUSTOMER

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var dto dtos.UpdateCustomerDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	customer := models.Customer{
		ID:               id,
		CustomerName:     dto.CustomerName,
		CustomerId:       dto.CustomerId,
		IsBusiness:       dto.IsBusiness,
		Address:          dto.Address,
		PhoneNumbers:     dto.PhoneNumbers,
		CustomerState:    dto.CustomerState,
		Email:            dto.Email,
		LastName:         dto.LastName,
		IdentifierTypeID: dto.IdentifierTypeID,
	}

	err = cc.Service.UpdateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) GetCustomerByEmail(c *gin.Context) {
	permissionId := config.PERMISSION_GET_CUSTOMER_BY_EMAIL

	if !cc.Auth.CheckPermission(c, permissionId) {
		return
	}

	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	email := c.Param("email")

	customer, err := cc.Service.GetCustomerByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
