package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type NutritionalNeedsCalculationRepository interface {
	CreateNutritionalNeedsCalculation(a *models.NutritionalNeedsCalculation) error
	GetNutritionalNeedsCalculationByID(id string) (*models.NutritionalNeedsCalculation, error)
	UpdateNutritionalNeedsCalculation(a *models.NutritionalNeedsCalculation) error
	DeleteNutritionalNeedsCalculation(id string) error
	ListNutritionalNeedsCalculations() ([]models.NutritionalNeedsCalculation, error)
}

type nutritionalNeedsCalculationRepo struct {
	db *gorm.DB
}

func NewNutritionalNeedsCalculationRepository(db *gorm.DB) NutritionalNeedsCalculationRepository {
	return &nutritionalNeedsCalculationRepo{db: db}
}

func (r *nutritionalNeedsCalculationRepo) CreateNutritionalNeedsCalculation(a *models.NutritionalNeedsCalculation) error {
	return r.db.Create(a).Error
}

func (r *nutritionalNeedsCalculationRepo) GetNutritionalNeedsCalculationByID(id string) (*models.NutritionalNeedsCalculation, error) {
	var a models.NutritionalNeedsCalculation
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *nutritionalNeedsCalculationRepo) UpdateNutritionalNeedsCalculation(a *models.NutritionalNeedsCalculation) error {
	return r.db.Save(a).Error
}

func (r *nutritionalNeedsCalculationRepo) DeleteNutritionalNeedsCalculation(id string) error {
	return r.db.Delete(&models.NutritionalNeedsCalculation{}, "id = ?", id).Error
}

func (r *nutritionalNeedsCalculationRepo) ListNutritionalNeedsCalculations() ([]models.NutritionalNeedsCalculation, error) {
	var items []models.NutritionalNeedsCalculation
	err := r.db.Find(&items).Error
	return items, err
}
