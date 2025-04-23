package services

import (
	"errors"
	"time"
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
	count, err := s.Repo.CountAppointmentsAtDateTime(appointment.DateTime)
	if err != nil {
		return nil, err
	}

	if count >= 3 {
		return nil, errors.New("no hay mas citas disponibles en este horario :v")
	}

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

func (s *AppointmentService) GetAppointmentByCustomerIDAndDate(customerID int, dateTime time.Time) (*models.Appointment, error) {
	return s.Repo.GetAppointmentByCustomerIDAndDate(customerID, dateTime)
}

func (s *AppointmentService) DeleteAppointmentByID(id int) error {
	return s.Repo.DeleteAppointmentByID(id)
}

func (s *AppointmentService) GetHourlyAppointmentCount(date time.Time) ([]int, error) {
	if s.Repo == nil {
		return nil, errors.New("appointment repository is not initialized")
	}
	return s.Repo.CountAppointmentsByHourOnDate(date)
}
