package service

import (
    "renutri/internal/models"
    "renutri/internal/repository"
)

// PatientService defines business operations on patients.
type PatientService interface {
    CreatePatient(patient *models.Patient) error
    GetPatient(id string) (*models.Patient, error)
    UpdatePatient(patient *models.Patient) error
    DeletePatient(id string) error
    ListPatients() ([]models.Patient, error)
}

type patientService struct {
    repo repository.PatientRepository
}

// NewPatientService creates a PatientService with the given repository.
func NewPatientService(repo repository.PatientRepository) PatientService {
    return &patientService{repo: repo}
}

func (s *patientService) CreatePatient(patient *models.Patient) error {
    return s.repo.CreatePatient(patient)
}

func (s *patientService) GetPatient(id string) (*models.Patient, error) {
    return s.repo.GetPatientByID(id)
}

func (s *patientService) UpdatePatient(patient *models.Patient) error {
    return s.repo.UpdatePatient(patient)
}

func (s *patientService) DeletePatient(id string) error {
    return s.repo.DeletePatient(id)
}

func (s *patientService) ListPatients() ([]models.Patient, error) {
    return s.repo.ListPatients()
}