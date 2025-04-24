package models

import "time"

// GeneratedDocument links to produced files (PDFs, etc.).
type GeneratedDocument struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	PatientID    string    `json:"patientId" gorm:"index;not null"`
	DocumentDate time.Time `json:"documentDate"`
	DocumentType string    `json:"documentType"`
	FilePath     string    `json:"filePath"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
