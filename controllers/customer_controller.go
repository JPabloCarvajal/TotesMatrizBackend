package controllers

import (
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
	Log     *utilities.LogUtil
}

func NewCustomerController(service *services.CustomerService, auth *utilities.AuthorizationUtil,
	log *utilities.LogUtil) *CustomerController {
	return &CustomerController{Service: service, Auth: auth, Log: log}
}

func (cc *CustomerController) GetAllCustomers(c *gin.Context) {
	if cc.Log.RegisterLog(c, "Attempting to retrieve all customers") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ALL_CUSTOMERS
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for GetAllCustomers")
		return
	}

	customers, err := cc.Service.GetAllCustomers()
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving customers: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Successfully retrieved all customers")
	c.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) GetCustomerByID(c *gin.Context) {
	idParam := c.Param("id")
	if cc.Log.RegisterLog(c, "Attempting to retrieve customer with ID: "+idParam) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_CUSTOMER_BY_ID
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for GetCustomerByID")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Invalid customer ID provided: "+idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	customer, err := cc.Service.GetCustomerByID(id)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Customer not found with ID: "+idParam)
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Successfully retrieved customer with ID: "+idParam)
	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) GetCustomerByCustomerID(c *gin.Context) {
	customerID := c.Param("customerID")
	if cc.Log.RegisterLog(c, "Attempting to retrieve customer with customerID: "+customerID) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_CUSTOMER_BY_CUSTOMERID
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for GetCustomerByCustomerID")
		return
	}

	customer, err := cc.Service.GetCustomerByCustomerID(customerID)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Customer not found with customerID: "+customerID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Successfully retrieved customer with customerID: "+customerID)
	c.JSON(http.StatusOK, customer)
}
func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	if cc.Log.RegisterLog(c, "Attempting to create new customer") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_CREATE_CUSTOMER
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for CreateCustomer")
		return
	}

	var dto dtos.CreateCustomerDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = cc.Log.RegisterLog(c, "Invalid JSON format in CreateCustomer request")
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
		_ = cc.Log.RegisterLog(c, "Error creating customer: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating customer"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Customer created successfully with CustomerID: "+createdCustomer.CustomerId)
	c.JSON(http.StatusCreated, createdCustomer)
}
func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
	if cc.Log.RegisterLog(c, "Attempting to update customer") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_UPDATE_CUSTOMER
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for UpdateCustomer")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Invalid customer ID format in URL parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var dto dtos.UpdateCustomerDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = cc.Log.RegisterLog(c, "Invalid JSON format in UpdateCustomer request")
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
		_ = cc.Log.RegisterLog(c, "Error updating customer with ID "+strconv.Itoa(id)+": "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating customer"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Customer updated successfully with ID: "+strconv.Itoa(id))
	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) GetCustomerByEmail(c *gin.Context) {
	if cc.Log.RegisterLog(c, "Attempting to retrieve customer by email") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_CUSTOMER_BY_EMAIL
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for GetCustomerByEmail")
		return
	}

	email := c.Param("email")

	customer, err := cc.Service.GetCustomerByEmail(email)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Customer not found with email: "+email)
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	_ = cc.Log.RegisterLog(c, "Customer retrieved successfully with email: "+email)
	c.JSON(http.StatusOK, customer)
}
func (cc *CustomerController) SearchCustomersByID(c *gin.Context) {
	if cc.Log.RegisterLog(c, "Attempting to search customers by ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_CUSTOMERS_BY_ID
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for SearchCustomersByID")
		return
	}

	query := c.Query("id")

	customers, err := cc.Service.SearchCustomersByID(query)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving customers by ID query: "+query)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}

	if len(customers) == 0 {
		_ = cc.Log.RegisterLog(c, "No customers found for ID query: "+query)
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

	_ = cc.Log.RegisterLog(c, "Customers retrieved successfully for ID query: "+query)
	c.JSON(http.StatusOK, customersDTO)
}

func (cc *CustomerController) SearchCustomersByName(c *gin.Context) {
	if cc.Log.RegisterLog(c, "Attempting to search customers by name") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_CUSTOMERS_BY_NAME
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for SearchCustomersByName")
		return
	}

	query := c.Query("name")

	customers, err := cc.Service.SearchCustomersByName(query)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving customers by name query: "+query)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}

	if len(customers) == 0 {
		_ = cc.Log.RegisterLog(c, "No customers found for name query: "+query)
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

	_ = cc.Log.RegisterLog(c, "Customers retrieved successfully for name query: "+query)
	c.JSON(http.StatusOK, customersDTO)
}

func (cc *CustomerController) SearchCustomersByLastName(c *gin.Context) {
	if cc.Log.RegisterLog(c, "Attempting to search customers by last name") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_CUSTOMERS_BY_LASTNAME
	if !cc.Auth.CheckPermission(c, permissionId) {
		_ = cc.Log.RegisterLog(c, "Access denied for SearchCustomersByLastName")
		return
	}

	query := c.Query("lastName")

	customers, err := cc.Service.SearchCustomersByLastName(query)
	if err != nil {
		_ = cc.Log.RegisterLog(c, "Error retrieving customers by last name query: "+query)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving customers"})
		return
	}

	if len(customers) == 0 {
		_ = cc.Log.RegisterLog(c, "No customers found for last name query: "+query)
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

	_ = cc.Log.RegisterLog(c, "Customers retrieved successfully for last name query: "+query)
	c.JSON(http.StatusOK, customersDTO)
}
