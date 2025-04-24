package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type AnthropometricAssessmentRepository interface {
	CreateAnthropometricAssessment(a *models.AnthropometricAssessment) error
	GetAnthropometricAssessmentByID(id string) (*models.AnthropometricAssessment, error)
	UpdateAnthropometricAssessment(a *models.AnthropometricAssessment) error
	DeleteAnthropometricAssessment(id string) error
	ListAnthropometricAssessments() ([]models.AnthropometricAssessment, error)
}

type anthropometricAssessmentRepo struct {
	db *gorm.DB
}

func NewAnthropometricAssessmentRepository(db *gorm.DB) AnthropometricAssessmentRepository {
	return &anthropometricAssessmentRepo{db: db}
}

func (r *anthropometricAssessmentRepo) CreateAnthropometricAssessment(a *models.AnthropometricAssessment) error {
	return r.db.Create(a).Error
}

func (r *anthropometricAssessmentRepo) GetAnthropometricAssessmentByID(id string) (*models.AnthropometricAssessment, error) {
	var a models.AnthropometricAssessment
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *anthropometricAssessmentRepo) UpdateAnthropometricAssessment(a *models.AnthropometricAssessment) error {
	return r.db.Save(a).Error
}

func (r *anthropometricAssessmentRepo) DeleteAnthropometricAssessment(id string) error {
	return r.db.Delete(&models.AnthropometricAssessment{}, "id = ?", id).Error
}

func (r *anthropometricAssessmentRepo) ListAnthropometricAssessments() ([]models.AnthropometricAssessment, error) {
	var assessments []models.AnthropometricAssessment
	err := r.db.Find(&assessments).Error
	return assessments, err
}
