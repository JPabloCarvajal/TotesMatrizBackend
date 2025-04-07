package controllers

import (
	"errors"
	"net/http"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	Service *services.EmployeeService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewEmployeeController(service *services.EmployeeService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *EmployeeController {
	return &EmployeeController{Service: service, Auth: auth, Log: log}
}

func (ec *EmployeeController) GetEmployeeByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_EMPLOYEE_BY_ID

	if ec.Log.RegisterLog(c, "Attempting to get employee by ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !ec.Auth.CheckPermission(c, permissionId) {
		_ = ec.Log.RegisterLog(c, "Access denied for GetEmployeeByID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	id := c.Param("id")

	employee, err := ec.Service.GetEmployeeByID(id)
	if err != nil {
		_ = ec.Log.RegisterLog(c, "Employee not found with ID: "+id)
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

	_ = ec.Log.RegisterLog(c, "Successfully retrieved employee with ID: "+id)
	c.JSON(http.StatusOK, employeeDTO)
}

func (ec *EmployeeController) GetAllEmployees(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_EMPLOYEES

	if ec.Log.RegisterLog(c, "Attempting to get all employees") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !ec.Auth.CheckPermission(c, permissionId) {
		_ = ec.Log.RegisterLog(c, "Access denied for GetAllEmployees")
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	employees, err := ec.Service.GetAllEmployees()
	if err != nil {
		_ = ec.Log.RegisterLog(c, "Error retrieving employees: "+err.Error())
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

	_ = ec.Log.RegisterLog(c, "Successfully retrieved all employees")
	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) SearchEmployeesByID(c *gin.Context) {
	query := c.Query("id")
	permissionId := config.PERMISSION_SEARCH_EMPLOYEES_BY_ID

	if ec.Log.RegisterLog(c, "Attempting to search employees by ID: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !ec.Auth.CheckPermission(c, permissionId) {
		_ = ec.Log.RegisterLog(c, "Access denied for SearchEmployeesByID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	employees, err := ec.Service.SearchEmployeesByID(query)
	if err != nil {
		_ = ec.Log.RegisterLog(c, "Error retrieving employees by ID: "+query+" - "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	if len(employees) == 0 {
		_ = ec.Log.RegisterLog(c, "No employees found with ID: "+query)
		c.JSON(http.StatusNotFound, gin.H{"message": "No employees found"})
		return
	}

	var employeesDTO []dtos.GetEmployeeDTO
	for _, employee := range employees {
		employeesDTO = append(employeesDTO, dtos.GetEmployeeDTO{
			ID:               employee.ID,
			Names:            employee.Names,
			LastNames:        employee.LastNames,
			PhoneNumbers:     employee.PhoneNumbers,
			UserID:           employee.UserID,
			IdentifierTypeID: employee.IdentifierTypeID,
		})
	}

	_ = ec.Log.RegisterLog(c, "Successfully found employees matching ID: "+query)
	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) SearchEmployeesByName(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_EMPLOYEES_BY_NAME

	query := c.Query("names")

	if ec.Log.RegisterLog(c, "Attempting to search employees by name: "+query) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !ec.Auth.CheckPermission(c, permissionId) {
		_ = ec.Log.RegisterLog(c, "Access denied for SearchEmployeesByName")
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if query == "" {
		_ = ec.Log.RegisterLog(c, "Empty name query provided in SearchEmployeesByName")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	employees, err := ec.Service.SearchEmployeesByName(query)
	if err != nil {
		_ = ec.Log.RegisterLog(c, "Error retrieving employees by name: "+query+" - "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	if len(employees) == 0 {
		_ = ec.Log.RegisterLog(c, "No employees found with name: "+query)
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

	_ = ec.Log.RegisterLog(c, "Successfully found employees with name: "+query)
	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_EMPLOYEE

	// Log de intento
	if err := ec.Log.RegisterLog(c, "Attempting to access CreateEmployee"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	// Verificar permiso
	if !ec.Auth.CheckPermission(c, permissionId) {
		_ = ec.Log.RegisterLog(c, "Permission denied for CreateEmployee")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	// Obtener empleados
	employees, err := ec.Service.GetAllEmployees()
	if err != nil {
		_ = ec.Log.RegisterLog(c, "Error retrieving employees: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	// Construir respuesta DTO
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

	_ = ec.Log.RegisterLog(c, "Employees retrieved successfully in CreateEmployee")
	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_EMPLOYEE

	if err := ec.Log.RegisterLog(c, "Attempting to update an employee"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !ec.Auth.CheckPermission(c, permissionId) {
		_ = ec.Log.RegisterLog(c, "Permission denied for UpdateEmployee")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	id := c.Param("id")

	var dto dtos.UpdateEmployeeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = ec.Log.RegisterLog(c, "Invalid JSON in UpdateEmployee: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := ec.Service.GetEmployeeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = ec.Log.RegisterLog(c, "Employee not found in UpdateEmployee: ID = "+id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		_ = ec.Log.RegisterLog(c, "Error retrieving employee in UpdateEmployee: "+err.Error())
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
		_ = ec.Log.RegisterLog(c, "Error updating employee: "+err.Error())
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

	_ = ec.Log.RegisterLog(c, "Successfully updated Employee with ID: "+id)

	c.JSON(http.StatusOK, employeeDTO)
}
