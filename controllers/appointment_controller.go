package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"totesbackend/controllers/utilities"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppointmentController struct {
	Service *services.AppointmentService
	Auth    *utilities.AuthorizationUtil
}

func NewAppointmentController(service *services.AppointmentService, auth *utilities.AuthorizationUtil) *AppointmentController {
	return &AppointmentController{Service: service, Auth: auth}
}

func (ac *AppointmentController) GetAppointmentByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	appointment, err := ac.Service.GetAppointmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ac *AppointmentController) GetAllAppointments(c *gin.Context) {

	appointments, err := ac.Service.GetAllAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) SearchAppointmentsByID(c *gin.Context) {

	query := c.Query("id")
	fmt.Println("Searching appointments by ID with:", query)

	appointments, err := ac.Service.SearchAppointmentsByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	if len(appointments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No appointments found"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) SearchAppointmentsByCustomerID(c *gin.Context) {

	query := c.Query("id")
	fmt.Println("Searching appointments by Customer ID with:", query)

	appointments, err := ac.Service.SearchAppointmentsByCustomerID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	if len(appointments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No appointments found"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) SearchAppointmentsByState(c *gin.Context) {

	state, err := strconv.ParseBool(c.Query("state"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state value"})
		return
	}

	appointments, err := ac.Service.SearchAppointmentsByState(state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) GetAppointmentsByCustomerID(c *gin.Context) {

	customerID, err := strconv.Atoi(c.Param("customerID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	appointments, err := ac.Service.GetAppointmentsByCustomerID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ac *AppointmentController) CreateAppointment(c *gin.Context) {

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	createdAppointment, err := ac.Service.CreateAppointment(appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating appointment"})
		return
	}

	c.JSON(http.StatusCreated, createdAppointment)
}

func (ac *AppointmentController) UpdateAppointment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	appointment.ID = id

	err = ac.Service.UpdateAppointment(&appointment)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating appointment"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ac *AppointmentController) GetAppointmentByCustomerIDAndDate(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Query("customerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	dateTime, err := time.Parse("2006-01-02 15:04:05", c.Query("dateTime"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use 'YYYY-MM-DD HH:MM:SS'"})
		return
	}

	appointment, err := ac.Service.GetAppointmentByCustomerIDAndDate(customerID, dateTime)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Filtrar solo los campos requeridos en la respuesta
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
