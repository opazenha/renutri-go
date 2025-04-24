package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type GuidanceService struct {
	repo repository.GuidanceRepository
}

func NewGuidanceService(repo repository.GuidanceRepository) *GuidanceService {
	return &GuidanceService{repo: repo}
}

func (s *GuidanceService) Create(a *models.Guidance) error {
	return s.repo.CreateGuidance(a)
}

func (s *GuidanceService) GetByID(id string) (*models.Guidance, error) {
	return s.repo.GetGuidanceByID(id)
}

func (s *GuidanceService) Update(a *models.Guidance) error {
	return s.repo.UpdateGuidance(a)
}

func (s *GuidanceService) Delete(id string) error {
	return s.repo.DeleteGuidance(id)
}

func (s *GuidanceService) List() ([]models.Guidance, error) {
	return s.repo.ListGuidances()
}
