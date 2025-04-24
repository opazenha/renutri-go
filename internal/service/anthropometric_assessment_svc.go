package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type AnthropometricAssessmentService struct {
	repo repository.AnthropometricAssessmentRepository
}

func NewAnthropometricAssessmentService(repo repository.AnthropometricAssessmentRepository) *AnthropometricAssessmentService {
	return &AnthropometricAssessmentService{repo: repo}
}

func (s *AnthropometricAssessmentService) Create(a *models.AnthropometricAssessment) error {
	return s.repo.CreateAnthropometricAssessment(a)
}

func (s *AnthropometricAssessmentService) GetByID(id string) (*models.AnthropometricAssessment, error) {
	return s.repo.GetAnthropometricAssessmentByID(id)
}

func (s *AnthropometricAssessmentService) Update(a *models.AnthropometricAssessment) error {
	return s.repo.UpdateAnthropometricAssessment(a)
}

func (s *AnthropometricAssessmentService) Delete(id string) error {
	return s.repo.DeleteAnthropometricAssessment(id)
}

func (s *AnthropometricAssessmentService) List() ([]models.AnthropometricAssessment, error) {
	return s.repo.ListAnthropometricAssessments()
}
