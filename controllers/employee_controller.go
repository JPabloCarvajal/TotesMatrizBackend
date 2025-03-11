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

type EmployeeController struct {
	Service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{Service: service}
}

func (ec *EmployeeController) GetEmployeeByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	employee, err := ec.Service.SearchEmployeeByID(id)
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
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	employees, err := ec.Service.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	var employeesDTO []dtos.GetEmployeeDTO
	for _, employee := range employees {
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
		employeesDTO = append(employeesDTO, employeeDTO)
	}

	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) SearchEmployeesByName(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

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
		employeesDTO = append(employeesDTO, employeeDTO)
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

	existingEmployee, _ := ec.Service.SearchEmployeeByID(dto.PersonalID)
	if existingEmployee != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "An employee with this Personal ID already exists"})
		return
	}

	if dto.UserID <= 0 || dto.IdentifierTypeID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID or Identifier Type ID"})
		return
	}

	employee := models.Employee{
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
	id := c.Param("id")

	var dto dtos.UpdateEmployeeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	employee, err := ec.Service.UpdateEmployee(id, dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating employee"})
		return
	}

	employeeDTO := dtos.UpdateEmployeeDTO{
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
