package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Appointment holds details about a single patient visit/interaction.
type Appointment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID     primitive.ObjectID `bson:"patient_id" json:"patientId"`
	Date          time.Time     `json:"date,omitempty"`
	Time          string        `json:"time,omitempty"`
	ServiceType   ServiceType   `json:"serviceType,omitempty"`
	Value         float64       `json:"value,omitempty"`
	PaymentMethod PaymentMethod `json:"paymentMethod,omitempty"`
	Notes         string        `json:"notes,omitempty"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}
