package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppointmentController struct {
	Service *services.AppointmentService
	Auth    *utilities.AuthorizationUtil
	Log     *utilities.LogUtil
}

func NewAppointmentController(service *services.AppointmentService, auth *utilities.AuthorizationUtil,
	log *utilities.LogUtil) *AppointmentController {
	return &AppointmentController{Service: service, Auth: auth, Log: log}
}

func (ac *AppointmentController) GetAppointmentByID(c *gin.Context) {
	if ac.Log.RegisterLog(c, "Attempting to get appointment by ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_APPOINTMENT_BY_ID
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for GetAppointmentByID")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid appointment ID: "+c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	appointment, err := ac.Service.GetAppointmentByID(id)
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Appointment not found for ID: "+strconv.Itoa(id))
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointment retrieved successfully for ID: "+strconv.Itoa(id))
	c.JSON(http.StatusOK, appointment)
}

func (ac *AppointmentController) GetAllAppointments(c *gin.Context) {
	if ac.Log.RegisterLog(c, "Attempting to retrieve all appointments") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_ALL_APPOINTMENTS
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for GetAllAppointments")
		return
	}

	appointments, err := ac.Service.GetAllAppointments()
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Error retrieving appointments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	_ = ac.Log.RegisterLog(c, "All appointments retrieved successfully")
	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) SearchAppointmentsByID(c *gin.Context) {

	if ac.Log.RegisterLog(c, "Attempting to search appointments by ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_APPOINTMENTS_BY_ID
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for SearchAppointmentsByID")
		return
	}

	query := c.Query("id")
	fmt.Println("Searching appointments by ID with:", query)

	appointments, err := ac.Service.SearchAppointmentsByID(query)
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Error retrieving appointments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	if len(appointments) == 0 {
		_ = ac.Log.RegisterLog(c, "No appointments found for given ID")
		c.JSON(http.StatusNotFound, gin.H{"message": "No appointments found"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointments found by ID successfully")
	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) SearchAppointmentsByCustomerID(c *gin.Context) {

	if ac.Log.RegisterLog(c, "Attempting to search appointments by customer ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_APPOINTMENT_BY_CUSTOMER_ID
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for SearchAppointmentsByCustomerID")
		return
	}

	query := c.Query("id")
	fmt.Println("Searching appointments by Customer ID with:", query)

	appointments, err := ac.Service.SearchAppointmentsByCustomerID(query)
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Error retrieving appointments by customer ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	if len(appointments) == 0 {
		_ = ac.Log.RegisterLog(c, "No appointments found for given customer ID")
		c.JSON(http.StatusNotFound, gin.H{"message": "No appointments found"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointments found by customer ID successfully")
	c.JSON(http.StatusOK, appointments)
}
func (ac *AppointmentController) SearchAppointmentsByState(c *gin.Context) {

	if ac.Log.RegisterLog(c, "Attempting to search appointments by state") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_SEARCH_APPOINTMENT_BY_STATE
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for SearchAppointmentsByState")
		return
	}

	state, err := strconv.ParseBool(c.Query("state"))
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid state value provided for appointment search")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state value"})
		return
	}

	appointments, err := ac.Service.SearchAppointmentsByState(state)
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Error retrieving appointments by state")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointments retrieved successfully by state")
	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) GetAppointmentsByCustomerID(c *gin.Context) {

	if ac.Log.RegisterLog(c, "Attempting to retrieve appointments by customer ID") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_APPOINTMENT_BY_CUSTOMER_ID
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for GetAppointmentsByCustomerID")
		return
	}

	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid customer ID provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	appointments, err := ac.Service.GetAppointmentsByCustomerID(customerID)
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Error retrieving appointments by customer ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointments retrieved successfully by customer ID")
	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) CreateAppointment(c *gin.Context) {
	if ac.Log.RegisterLog(c, "Attempting to create appointment") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_CREATE_APPOINTMENT
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for CreateAppointment")
		c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para crear citas"})
		return
	}

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid JSON format when creating appointment")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
		return
	}

	createdAppointment, err := ac.Service.CreateAppointment(appointment)
	if err != nil {
		if err.Error() == "ya existen 3 citas agendadas para esta fecha y hora" {
			_ = ac.Log.RegisterLog(c, "limite de citas alcanzado :v")
			c.JSON(http.StatusBadRequest, gin.H{"error": "no se puede crear la cita. Ya hay 3 citas agendadas para esta fecha y hora."})
		} else {
			_ = ac.Log.RegisterLog(c, "Error creando cita")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la cita"})
		}
		return
	}

	_ = ac.Log.RegisterLog(c, "Cita creada exitosamente")
	c.JSON(http.StatusCreated, createdAppointment)
}

func (ac *AppointmentController) UpdateAppointment(c *gin.Context) {

	if ac.Log.RegisterLog(c, "Attempting to update appointment") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_UPDATE_APPOINTMENT

	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for UpdateAppointment")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid appointment ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid JSON format on update appointment")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	appointment.ID = id

	err = ac.Service.UpdateAppointment(&appointment)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = ac.Log.RegisterLog(c, "Appointment not found for update")
			c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
			return
		}
		_ = ac.Log.RegisterLog(c, "Error updating appointment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating appointment"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointment updated successfully")
	c.JSON(http.StatusOK, appointment)
}

func (ac *AppointmentController) GetAppointmentByCustomerIDAndDate(c *gin.Context) {

	if ac.Log.RegisterLog(c, "Attempting to get appointment by customer ID and date") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_GET_APPOINTMENTS_BY_CUSTOMERID_AND_DATE

	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for GetAppointmentByCustomerIDAndDate")
		return
	}

	customerID, err := strconv.Atoi(c.Query("customerId"))
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid customer ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	dateTime, err := time.Parse("2006-01-02 15:04:05", c.Query("dateTime"))
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid date format for GetAppointmentByCustomerIDAndDate")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use 'YYYY-MM-DD HH:MM:SS'"})
		return
	}

	appointment, err := ac.Service.GetAppointmentByCustomerIDAndDate(customerID, dateTime)
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Appointment not found for given customer ID and date")
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointment successfully retrieved by customer ID and date")

	response := gin.H{
		"id":           appointment.ID,
		"dateTime":     appointment.DateTime,
		"customerName": appointment.CustomerName,
		"lastName":     appointment.LastName,
		"email":        appointment.Email,
		"customerID":   appointment.CustomerID,
	}

	c.JSON(http.StatusOK, response)
}

func (ac *AppointmentController) DeleteAppointmentByID(c *gin.Context) {
	if ac.Log.RegisterLog(c, "Attempting to delete appointment") != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	permissionId := config.PERMISSION_DELETE_APPOINTMENT
	if !ac.Auth.CheckPermission(c, permissionId) {
		_ = ac.Log.RegisterLog(c, "Access denied for DeleteAppointmentByID")
		c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para eliminar citas"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = ac.Log.RegisterLog(c, "Invalid appointment ID: "+c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	err = ac.Service.DeleteAppointmentByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = ac.Log.RegisterLog(c, "Appointment not found for ID: "+strconv.Itoa(id))
			c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		} else {
			_ = ac.Log.RegisterLog(c, "Error deleting appointment")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	_ = ac.Log.RegisterLog(c, "Appointment deleted successfully for ID: "+strconv.Itoa(id))
	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
