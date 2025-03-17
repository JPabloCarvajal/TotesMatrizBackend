package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type AppointmentService struct {
	Repo *repositories.AppointmentRepository
}

func NewAppointmentService(repo *repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{Repo: repo}
}

func (s *AppointmentService) GetAppointmentByID(id int) (*models.Appointment, error) {
	return s.Repo.GetAppointmentByID(id)
}

func (s *AppointmentService) GetAllAppointments() ([]models.Appointment, error) {
	return s.Repo.GetAllAppointments()
}

func (s *AppointmentService) SearchAppointmentsByState(state bool) ([]models.Appointment, error) {
	return s.Repo.SearchAppointmentsByState(state)
}

func (s *AppointmentService) GetAppointmentsByCustomerID(customerID int) ([]models.Appointment, error) {
	return s.Repo.GetAppointmentsByCustomerID(customerID)
}

func (s *AppointmentService) CreateAppointment(appointment models.Appointment) (*models.Appointment, error) {
	return s.Repo.CreateAppointment(&appointment)
}

func (s *AppointmentService) UpdateAppointment(appointment *models.Appointment) error {
	return s.Repo.UpdateAppointment(appointment)
}

func (s *AppointmentService) SearchAppointmentsByID(query string) ([]models.Appointment, error) {
	return s.Repo.SearchAppointmentsByID(query)
}

func (s *AppointmentService) SearchAppointmentsByCustomerID(query string) ([]models.Appointment, error) {
	return s.Repo.SearchAppointmentsByCustomerID(query)
}
