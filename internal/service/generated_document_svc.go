package service

import (
	"renutri/internal/models"
	"renutri/internal/repository"
)

type GeneratedDocumentService struct {
	repo repository.GeneratedDocumentRepository
}

func NewGeneratedDocumentService(repo repository.GeneratedDocumentRepository) *GeneratedDocumentService {
	return &GeneratedDocumentService{repo: repo}
}

func (s *GeneratedDocumentService) Create(a *models.GeneratedDocument) error {
	return s.repo.CreateGeneratedDocument(a)
}

func (s *GeneratedDocumentService) GetByID(id string) (*models.GeneratedDocument, error) {
	return s.repo.GetGeneratedDocumentByID(id)
}

func (s *GeneratedDocumentService) Update(a *models.GeneratedDocument) error {
	return s.repo.UpdateGeneratedDocument(a)
}

func (s *GeneratedDocumentService) Delete(id string) error {
	return s.repo.DeleteGeneratedDocument(id)
}

func (s *GeneratedDocumentService) List() ([]models.GeneratedDocument, error) {
	return s.repo.ListGeneratedDocuments()
}
