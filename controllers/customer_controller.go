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

func (cc *CustomerController) GetCustomers(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	customers, err := cc.Service.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) GetCustomerByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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

func (cc *CustomerController) SearchCustomersByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("id")
	fmt.Println("Searching customers by ID with:", query)

	customers, err := cc.Service.SearchCustomersByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}

	if len(customers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No customers found"})
		return
	}

	var customersDTO []dtos.GetCustomerDTO
	for _, customer := range customers {
		customersDTO = append(customersDTO, dtos.GetCustomerDTO{
			ID:               customer.ID,
			CustomerName:     customer.CustomerName,
			CustomerId:       customer.CustomerId,
			IsBusiness:       customer.IsBusiness,
			Address:          customer.Address,
			PhoneNumbers:     customer.PhoneNumbers,
			CustomerState:    customer.CustomerState,
			Email:            customer.Email,
			LastName:         customer.LastName,
			IdentifierTypeID: customer.IdentifierTypeID,
		})
	}

	c.JSON(http.StatusOK, customersDTO)
}

func (cc *CustomerController) SearchCustomersByName(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("name")
	fmt.Println("Searching customers by name with:", query)

	customers, err := cc.Service.SearchCustomersByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}

	if len(customers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No customers found"})
		return
	}

	var customersDTO []dtos.GetCustomerDTO
	for _, customer := range customers {
		customersDTO = append(customersDTO, dtos.GetCustomerDTO{
			ID:               customer.ID,
			CustomerName:     customer.CustomerName,
			CustomerId:       customer.CustomerId,
			IsBusiness:       customer.IsBusiness,
			Address:          customer.Address,
			PhoneNumbers:     customer.PhoneNumbers,
			CustomerState:    customer.CustomerState,
			Email:            customer.Email,
			LastName:         customer.LastName,
			IdentifierTypeID: customer.IdentifierTypeID,
		})
	}

	c.JSON(http.StatusOK, customersDTO)
}
