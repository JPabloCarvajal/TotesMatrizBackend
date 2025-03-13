package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type AppointmentRepository struct {
	DB *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{DB: db}
}

func (r *AppointmentRepository) GetAppointmentByID(id int) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.DB.First(&appointment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *AppointmentRepository) GetAllAppointments() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) SearchAppointmentsByState(state bool) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Where("state = ?", state).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetAppointmentsByCustomerID(customerID int) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Where("customer_id = ?", customerID).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) CreateAppointment(appointment *models.Appointment) (*models.Appointment, error) {
	if err := r.DB.Create(appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (r *AppointmentRepository) UpdateAppointment(appointment *models.Appointment) error {
	if err := r.DB.Save(appointment).Error; err != nil {
		return err
	}
	return nil
}
