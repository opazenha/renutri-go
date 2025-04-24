package models

import "time"

// Guidance captures nutritionist advice or instructions.
type Guidance struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	PatientID string    `json:"patientId" gorm:"index;not null"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
