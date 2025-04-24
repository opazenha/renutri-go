package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

// AppointmentService defines business logic for appointments.
type AppointmentService interface {
	CreateAppointment(app *models.Appointment) error
	GetAppointment(id string) (*models.Appointment, error)
	UpdateAppointment(app *models.Appointment) error
	DeleteAppointment(id string) error
	ListAppointments() ([]models.Appointment, error)
	ListByPatient(patientID string) ([]models.Appointment, error)
}

type appointmentService struct {
	repo repository.AppointmentRepository
}

// NewAppointmentService constructs an AppointmentService.
func NewAppointmentService(repo repository.AppointmentRepository) AppointmentService {
	return &appointmentService{repo: repo}
}

func (s *appointmentService) CreateAppointment(app *models.Appointment) error {
	return s.repo.CreateAppointment(app)
}

func (s *appointmentService) GetAppointment(id string) (*models.Appointment, error) {
	return s.repo.GetAppointmentByID(id)
}

func (s *appointmentService) UpdateAppointment(app *models.Appointment) error {
	return s.repo.UpdateAppointment(app)
}

func (s *appointmentService) DeleteAppointment(id string) error {
	return s.repo.DeleteAppointment(id)
}

func (s *appointmentService) ListAppointments() ([]models.Appointment, error) {
	return s.repo.ListAppointments()
}

func (s *appointmentService) ListByPatient(patientID string) ([]models.Appointment, error) {
	return s.repo.ListByPatient(patientID)
}
