package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Service *services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

// Obtener todos los clientes
func (cc *CustomerController) GetCustomers(c *gin.Context) {
	customers, err := cc.Service.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// Obtener cliente por ID
func (cc *CustomerController) GetCustomerByID(c *gin.Context) {
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

// Crear un cliente
func (cc *CustomerController) CreateCustomer(c *gin.Context) {
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

// Actualizar un cliente
func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
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
