package models

import (
	"time"
)

// HealthCondition records a medical condition for a patient.
type HealthCondition struct {
	ID            string     `json:"id" gorm:"primaryKey"`
	PatientID     string     `json:"patientId" gorm:"index;not null"`
	ConditionName string     `json:"conditionName"`
	DiagnosisDate *time.Time `json:"diagnosisDate,omitempty"`
	Status        string     `json:"status"`
	Medication    string     `json:"medication,omitempty"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}
