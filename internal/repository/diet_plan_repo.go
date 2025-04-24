package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type DietPlanRepository interface {
	CreateDietPlan(a *models.DietPlan) error
	GetDietPlanByID(id string) (*models.DietPlan, error)
	UpdateDietPlan(a *models.DietPlan) error
	DeleteDietPlan(id string) error
	ListDietPlans() ([]models.DietPlan, error)
}

type dietPlanRepo struct {
	db *gorm.DB
}

func NewDietPlanRepository(db *gorm.DB) DietPlanRepository {
	return &dietPlanRepo{db: db}
}

func (r *dietPlanRepo) CreateDietPlan(a *models.DietPlan) error {
	return r.db.Create(a).Error
}

func (r *dietPlanRepo) GetDietPlanByID(id string) (*models.DietPlan, error) {
	var a models.DietPlan
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *dietPlanRepo) UpdateDietPlan(a *models.DietPlan) error {
	return r.db.Save(a).Error
}

func (r *dietPlanRepo) DeleteDietPlan(id string) error {
	return r.db.Delete(&models.DietPlan{}, "id = ?", id).Error
}

func (r *dietPlanRepo) ListDietPlans() ([]models.DietPlan, error) {
	var items []models.DietPlan
	err := r.db.Find(&items).Error
	return items, err
}
