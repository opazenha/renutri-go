package models

import "time"

// ClinicalEvaluation stores assessment date, symptoms, and lab results.
type ClinicalEvaluation struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	PatientID        string    `json:"patientId" gorm:"index;not null"`
	EvaluationDate   time.Time `json:"evaluationDate"`
	Symptoms         string    `json:"symptoms"`
	LaboratoryResults string   `json:"laboratoryResults"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
