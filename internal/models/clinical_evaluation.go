package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ClinicalEvaluation represents a clinical evaluation for MongoDB.
type ClinicalEvaluation struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID        primitive.ObjectID `bson:"patientId,omitempty" json:"patientId"`
	EvaluationDate   time.Time          `bson:"evaluationDate,omitempty" json:"evaluationDate"`
	Symptoms         string             `bson:"symptoms,omitempty" json:"symptoms"`
	LaboratoryResults string            `bson:"laboratoryResults,omitempty" json:"laboratoryResults"`
	CreatedAt        time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
