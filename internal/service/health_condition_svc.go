package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type HealthConditionService struct {
	repo repository.HealthConditionRepository
}

func NewHealthConditionService(repo repository.HealthConditionRepository) *HealthConditionService {
	return &HealthConditionService{repo: repo}
}

func (s *HealthConditionService) Create(a *models.HealthCondition) error {
	return s.repo.CreateHealthCondition(a)
}

func (s *HealthConditionService) GetByID(id string) (*models.HealthCondition, error) {
	return s.repo.GetHealthConditionByID(id)
}

func (s *HealthConditionService) Update(a *models.HealthCondition) error {
	return s.repo.UpdateHealthCondition(a)
}

func (s *HealthConditionService) Delete(id string) error {
	return s.repo.DeleteHealthCondition(id)
}

func (s *HealthConditionService) List() ([]models.HealthCondition, error) {
	return s.repo.ListHealthConditions()
}
