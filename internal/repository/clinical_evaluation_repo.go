package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type ClinicalEvaluationRepository interface {
	CreateClinicalEvaluation(a *models.ClinicalEvaluation) error
	GetClinicalEvaluationByID(id string) (*models.ClinicalEvaluation, error)
	UpdateClinicalEvaluation(a *models.ClinicalEvaluation) error
	DeleteClinicalEvaluation(id string) error
	ListClinicalEvaluations() ([]models.ClinicalEvaluation, error)
}

type clinicalEvaluationRepo struct {
	db *gorm.DB
}

func NewClinicalEvaluationRepository(db *gorm.DB) ClinicalEvaluationRepository {
	return &clinicalEvaluationRepo{db: db}
}

func (r *clinicalEvaluationRepo) CreateClinicalEvaluation(a *models.ClinicalEvaluation) error {
	return r.db.Create(a).Error
}

func (r *clinicalEvaluationRepo) GetClinicalEvaluationByID(id string) (*models.ClinicalEvaluation, error) {
	var a models.ClinicalEvaluation
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *clinicalEvaluationRepo) UpdateClinicalEvaluation(a *models.ClinicalEvaluation) error {
	return r.db.Save(a).Error
}

func (r *clinicalEvaluationRepo) DeleteClinicalEvaluation(id string) error {
	return r.db.Delete(&models.ClinicalEvaluation{}, "id = ?", id).Error
}

func (r *clinicalEvaluationRepo) ListClinicalEvaluations() ([]models.ClinicalEvaluation, error) {
	var items []models.ClinicalEvaluation
	err := r.db.Find(&items).Error
	return items, err
}
