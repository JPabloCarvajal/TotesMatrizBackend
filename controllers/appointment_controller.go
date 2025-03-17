package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppointmentController struct {
	Service *services.AppointmentService
}

func NewAppointmentController(service *services.AppointmentService) *AppointmentController {
	return &AppointmentController{Service: service}
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
