package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type ClinicalEvaluationService struct {
	repo repository.ClinicalEvaluationRepository
}

func NewClinicalEvaluationService(repo repository.ClinicalEvaluationRepository) *ClinicalEvaluationService {
	return &ClinicalEvaluationService{repo: repo}
}

func (s *ClinicalEvaluationService) Create(a *models.ClinicalEvaluation) error {
	return s.repo.CreateClinicalEvaluation(a)
}

func (s *ClinicalEvaluationService) GetByID(id string) (*models.ClinicalEvaluation, error) {
	return s.repo.GetClinicalEvaluationByID(id)
}

func (s *ClinicalEvaluationService) Update(a *models.ClinicalEvaluation) error {
	return s.repo.UpdateClinicalEvaluation(a)
}

func (s *ClinicalEvaluationService) Delete(id string) error {
	return s.repo.DeleteClinicalEvaluation(id)
}

func (s *ClinicalEvaluationService) List() ([]models.ClinicalEvaluation, error) {
	return s.repo.ListClinicalEvaluations()
}
