package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HealthCondition represents a patient's health condition for MongoDB.
type HealthCondition struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID     primitive.ObjectID `bson:"patientId,omitempty" json:"patientId,omitempty"`
	ConditionName string             `bson:"conditionName,omitempty" json:"conditionName"`
	DiagnosisDate *time.Time         `bson:"diagnosisDate,omitempty" json:"diagnosisDate,omitempty"`
	Status        string             `bson:"status,omitempty" json:"status"`
	Medication    string             `bson:"medication,omitempty" json:"medication,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
