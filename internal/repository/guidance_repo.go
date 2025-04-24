package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type GuidanceRepository interface {
	CreateGuidance(a *models.Guidance) error
	GetGuidanceByID(id string) (*models.Guidance, error)
	UpdateGuidance(a *models.Guidance) error
	DeleteGuidance(id string) error
	ListGuidances() ([]models.Guidance, error)
}

type guidanceRepo struct {
	db *gorm.DB
}

func NewGuidanceRepository(db *gorm.DB) GuidanceRepository {
	return &guidanceRepo{db: db}
}

func (r *guidanceRepo) CreateGuidance(a *models.Guidance) error {
	return r.db.Create(a).Error
}

func (r *guidanceRepo) GetGuidanceByID(id string) (*models.Guidance, error) {
	var a models.Guidance
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *guidanceRepo) UpdateGuidance(a *models.Guidance) error {
	return r.db.Save(a).Error
}

func (r *guidanceRepo) DeleteGuidance(id string) error {
	return r.db.Delete(&models.Guidance{}, "id = ?", id).Error
}

func (r *guidanceRepo) ListGuidances() ([]models.Guidance, error) {
	var items []models.Guidance
	err := r.db.Find(&items).Error
	return items, err
}
