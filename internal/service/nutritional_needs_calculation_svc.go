package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type NutritionalNeedsCalculationService struct {
	repo repository.NutritionalNeedsCalculationRepository
}

func NewNutritionalNeedsCalculationService(repo repository.NutritionalNeedsCalculationRepository) *NutritionalNeedsCalculationService {
	return &NutritionalNeedsCalculationService{repo: repo}
}

func (s *NutritionalNeedsCalculationService) Create(a *models.NutritionalNeedsCalculation) error {
	return s.repo.CreateNutritionalNeedsCalculation(a)
}

func (s *NutritionalNeedsCalculationService) GetByID(id string) (*models.NutritionalNeedsCalculation, error) {
	return s.repo.GetNutritionalNeedsCalculationByID(id)
}

func (s *NutritionalNeedsCalculationService) Update(a *models.NutritionalNeedsCalculation) error {
	return s.repo.UpdateNutritionalNeedsCalculation(a)
}

func (s *NutritionalNeedsCalculationService) Delete(id string) error {
	return s.repo.DeleteNutritionalNeedsCalculation(id)
}

func (s *NutritionalNeedsCalculationService) List() ([]models.NutritionalNeedsCalculation, error) {
	return s.repo.ListNutritionalNeedsCalculations()
}
