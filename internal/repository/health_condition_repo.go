package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type HealthConditionRepository interface {
	CreateHealthCondition(a *models.HealthCondition) error
	GetHealthConditionByID(id string) (*models.HealthCondition, error)
	UpdateHealthCondition(a *models.HealthCondition) error
	DeleteHealthCondition(id string) error
	ListHealthConditions() ([]models.HealthCondition, error)
}

type healthConditionRepo struct {
	db *gorm.DB
}

func NewHealthConditionRepository(db *gorm.DB) HealthConditionRepository {
	return &healthConditionRepo{db: db}
}

func (r *healthConditionRepo) CreateHealthCondition(a *models.HealthCondition) error {
	return r.db.Create(a).Error
}

func (r *healthConditionRepo) GetHealthConditionByID(id string) (*models.HealthCondition, error) {
	var a models.HealthCondition
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *healthConditionRepo) UpdateHealthCondition(a *models.HealthCondition) error {
	return r.db.Save(a).Error
}

func (r *healthConditionRepo) DeleteHealthCondition(id string) error {
	return r.db.Delete(&models.HealthCondition{}, "id = ?", id).Error
}

func (r *healthConditionRepo) ListHealthConditions() ([]models.HealthCondition, error) {
	var items []models.HealthCondition
	err := r.db.Find(&items).Error
	return items, err
}
