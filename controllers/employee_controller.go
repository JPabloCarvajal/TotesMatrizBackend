package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	Service *services.EmployeeService
	Auth    *utilities.AuthorizationUtil
}

func NewEmployeeController(service *services.EmployeeService, auth *utilities.AuthorizationUtil) *EmployeeController {
	return &EmployeeController{Service: service, Auth: auth}
}

func (ec *EmployeeController) GetEmployeeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_EMPLOYEE_BY_ID

	if !ec.Auth.CheckPermission(c, permissionId) {
		return
	}

	id := c.Param("id")

	employee, err := ec.Service.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	employeeDTO := dtos.GetEmployeeDTO{
		ID:               employee.ID,
		Names:            employee.Names,
		LastNames:        employee.LastNames,
		PersonalID:       employee.PersonalID,
		Address:          employee.Address,
		PhoneNumbers:     employee.PhoneNumbers,
		UserID:           employee.UserID,
		IdentifierTypeID: employee.IdentifierTypeID,
	}

	c.JSON(http.StatusOK, employeeDTO)
}


func (ec *EmployeeController) GetAllEmployees(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_EMPLOYEES

	if !ec.Auth.CheckPermission(c, permissionId) {
		return
	}


	employees, err := ec.Service.SearchEmployeesByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	if len(employees) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No employees found"})
		return
	}

	var employeesDTO []dtos.GetEmployeeDTO
	for _, employee := range employees {
		employeesDTO = append(employeesDTO, dtos.GetEmployeeDTO{
			ID:               employee.ID,
			Names:            employee.Names,
			LastNames:        employee.LastNames,
			PersonalID:       employee.PersonalID,
			Address:          employee.Address,
			PhoneNumbers:     employee.PhoneNumbers,
			UserID:           employee.UserID,
			IdentifierTypeID: employee.IdentifierTypeID,
		})
	}

	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) SearchEmployeesByName(c *gin.Context) {

	permissionId := config.PERMISSION_SEARCH_EMPLOYEES_BY_NAME

	if !ec.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("names")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}


	employees, err := ec.Service.SearchEmployeesByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	if len(employees) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No employees found"})
		return
	}

	var employeesDTO []dtos.GetEmployeeDTO
	for _, employee := range employees {
		employeesDTO = append(employeesDTO, dtos.GetEmployeeDTO{
			ID:               employee.ID,
			Names:            employee.Names,
			LastNames:        employee.LastNames,
			PersonalID:       employee.PersonalID,
			Address:          employee.Address,
			PhoneNumbers:     employee.PhoneNumbers,
			UserID:           employee.UserID,
			IdentifierTypeID: employee.IdentifierTypeID,
		})
	}

	c.JSON(http.StatusOK, employeesDTO)
}


func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_EMPLOYEE

	if !ec.Auth.CheckPermission(c, permissionId) {
		return
	}

	employees, err := ec.Service.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	var employeesDTO []dtos.GetEmployeeDTO
	for _, employee := range employees {
		employeesDTO = append(employeesDTO, dtos.GetEmployeeDTO{
			ID:               employee.ID,
			Names:            employee.Names,
			LastNames:        employee.LastNames,
			PersonalID:       employee.PersonalID,
			Address:          employee.Address,
			PhoneNumbers:     employee.PhoneNumbers,
			UserID:           employee.UserID,
			IdentifierTypeID: employee.IdentifierTypeID,
		})
	}

	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	var dto dtos.CreateEmployeeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	employee := &models.Employee{
		Names:            dto.Names,
		LastNames:        dto.LastNames,
		PersonalID:       dto.PersonalID,
		Address:          dto.Address,
		PhoneNumbers:     dto.PhoneNumbers,
		UserID:           dto.UserID,
		IdentifierTypeID: dto.IdentifierTypeID,
	}

	createdEmployee, err := ec.Service.CreateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating employee", "details": err.Error()})
		return
	}

	employeeDTO := dtos.GetEmployeeDTO{
		ID:               createdEmployee.ID,
		Names:            createdEmployee.Names,
		LastNames:        createdEmployee.LastNames,
		PersonalID:       createdEmployee.PersonalID,
		Address:          createdEmployee.Address,
		PhoneNumbers:     createdEmployee.PhoneNumbers,
		UserID:           createdEmployee.UserID,
		IdentifierTypeID: createdEmployee.IdentifierTypeID,
	}

	c.JSON(http.StatusCreated, employeeDTO)
}

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_EMPLOYEE

	if !ec.Auth.CheckPermission(c, permissionId) {
		return
	}
	id := c.Param("id")

	var dto dtos.UpdateEmployeeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := ec.Service.GetEmployeeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	employee.Names = dto.Names
	employee.LastNames = dto.LastNames
	employee.PersonalID = dto.PersonalID
	employee.Address = dto.Address
	employee.PhoneNumbers = dto.PhoneNumbers
	employee.UserID = dto.UserID
	employee.IdentifierTypeID = dto.IdentifierTypeID

	err = ec.Service.UpdateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	employeeDTO := dtos.GetEmployeeDTO{
		ID:               employee.ID,
		Names:            employee.Names,
		LastNames:        employee.LastNames,
		PersonalID:       employee.PersonalID,
		Address:          employee.Address,
		PhoneNumbers:     employee.PhoneNumbers,
		UserID:           employee.UserID,
		IdentifierTypeID: employee.IdentifierTypeID,
	}

	c.JSON(http.StatusOK, employeeDTO)
}
