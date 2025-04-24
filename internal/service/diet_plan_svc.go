package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type DietPlanService struct {
	repo repository.DietPlanRepository
}

func NewDietPlanService(repo repository.DietPlanRepository) *DietPlanService {
	return &DietPlanService{repo: repo}
}

func (s *DietPlanService) Create(a *models.DietPlan) error {
	return s.repo.CreateDietPlan(a)
}

func (s *DietPlanService) GetByID(id string) (*models.DietPlan, error) {
	return s.repo.GetDietPlanByID(id)
}

func (s *DietPlanService) Update(a *models.DietPlan) error {
	return s.repo.UpdateDietPlan(a)
}

func (s *DietPlanService) Delete(id string) error {
	return s.repo.DeleteDietPlan(id)
}

func (s *DietPlanService) List() ([]models.DietPlan, error) {
	return s.repo.ListDietPlans()
}
