package repository

import (
    "renutri/internal/models"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// PatientRepository defines CRUD operations for patients.
type PatientRepository interface {
    CreatePatient(patient *models.Patient) error
    GetPatientByID(id string) (*models.Patient, error)
    UpdatePatient(patient *models.Patient) error
    DeletePatient(id string) error
    ListPatients() ([]models.Patient, error)
}

type patientRepo struct {
    db *gorm.DB
}

// NewPatientRepository returns a new PatientRepository using GORM.
func NewPatientRepository(db *gorm.DB) PatientRepository {
    return &patientRepo{db}
}

func (r *patientRepo) CreatePatient(patient *models.Patient) error {
    // ensure ID is a valid UUID
    if patient.ID == "" {
        patient.ID = uuid.New().String()
    }
    return r.db.Create(patient).Error
}

func (r *patientRepo) GetPatientByID(id string) (*models.Patient, error) {
    var patient models.Patient
    err := r.db.First(&patient, "id = ?", id).Error
    return &patient, err
}

func (r *patientRepo) UpdatePatient(patient *models.Patient) error {
    return r.db.Save(patient).Error
}

func (r *patientRepo) DeletePatient(id string) error {
    return r.db.Delete(&models.Patient{}, "id = ?", id).Error
}

func (r *patientRepo) ListPatients() ([]models.Patient, error) {
    var patients []models.Patient
    err := r.db.Find(&patients).Error
    return patients, err
}