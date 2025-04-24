package models

import (
	"time"
)

// Appointment holds details about a single patient visit/interaction.
type Appointment struct {
	ID            string        `json:"id" gorm:"primaryKey"`
	PatientID     string        `json:"patientId" gorm:"index;not null"`
	Date          time.Time     `json:"date,omitempty"`
	Time          string        `json:"time,omitempty"`
	ServiceType   ServiceType   `json:"serviceType,omitempty"`
	Value         float64       `json:"value,omitempty"`
	PaymentMethod PaymentMethod `json:"paymentMethod,omitempty"`
	Notes         string        `json:"notes,omitempty"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}
