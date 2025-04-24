package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"renutri/internal/models"
)

// AppointmentRepository defines CRUD for appointment records.
type AppointmentRepository interface {
	CreateAppointment(app *models.Appointment) error
	GetAppointmentByID(id string) (*models.Appointment, error)
	UpdateAppointment(app *models.Appointment) error
	DeleteAppointment(id string) error
	ListAppointments() ([]models.Appointment, error)
	ListByPatient(patientID string) ([]models.Appointment, error)
}

type appointmentRepo struct {
	db *gorm.DB
}

// NewAppointmentRepository returns a GORM-based AppointmentRepository.
func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepo{db: db}
}

func (r *appointmentRepo) CreateAppointment(app *models.Appointment) error {
	if app.ID == "" {
		app.ID = uuid.New().String()
	}
	return r.db.Create(app).Error
}

func (r *appointmentRepo) GetAppointmentByID(id string) (*models.Appointment, error) {
	var app models.Appointment
	err := r.db.First(&app, "id = ?", id).Error
	return &app, err
}

func (r *appointmentRepo) UpdateAppointment(app *models.Appointment) error {
	return r.db.Save(app).Error
}

func (r *appointmentRepo) DeleteAppointment(id string) error {
	return r.db.Delete(&models.Appointment{}, "id = ?", id).Error
}

func (r *appointmentRepo) ListAppointments() ([]models.Appointment, error) {
	var apps []models.Appointment
	err := r.db.Find(&apps).Error
	return apps, err
}

func (r *appointmentRepo) ListByPatient(patientID string) ([]models.Appointment, error) {
	var apps []models.Appointment
	err := r.db.Where("patient_id = ?", patientID).Find(&apps).Error
	return apps, err
}
