package repository

import (
	"gorm.io/gorm"
	"renutri/internal/models"
)

type GeneratedDocumentRepository interface {
	CreateGeneratedDocument(a *models.GeneratedDocument) error
	GetGeneratedDocumentByID(id string) (*models.GeneratedDocument, error)
	UpdateGeneratedDocument(a *models.GeneratedDocument) error
	DeleteGeneratedDocument(id string) error
	ListGeneratedDocuments() ([]models.GeneratedDocument, error)
}

type generatedDocumentRepo struct {
	db *gorm.DB
}

func NewGeneratedDocumentRepository(db *gorm.DB) GeneratedDocumentRepository {
	return &generatedDocumentRepo{db: db}
}

func (r *generatedDocumentRepo) CreateGeneratedDocument(a *models.GeneratedDocument) error {
	return r.db.Create(a).Error
}

func (r *generatedDocumentRepo) GetGeneratedDocumentByID(id string) (*models.GeneratedDocument, error) {
	var a models.GeneratedDocument
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *generatedDocumentRepo) UpdateGeneratedDocument(a *models.GeneratedDocument) error {
	return r.db.Save(a).Error
}

func (r *generatedDocumentRepo) DeleteGeneratedDocument(id string) error {
	return r.db.Delete(&models.GeneratedDocument{}, "id = ?", id).Error
}

func (r *generatedDocumentRepo) ListGeneratedDocuments() ([]models.GeneratedDocument, error) {
	var items []models.GeneratedDocument
	err := r.db.Find(&items).Error
	return items, err
}
