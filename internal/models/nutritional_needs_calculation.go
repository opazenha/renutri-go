package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NutritionalNeedsCalculation records metabolic and nutrient requirements.
type NutritionalNeedsCalculation struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID       primitive.ObjectID `bson:"patientId,omitempty" json:"patientId"`
	CalculationDate time.Time          `bson:"calculationDate,omitempty" json:"calculationDate"`
	TMB             float64            `bson:"tmb,omitempty" json:"tmb"`
	GEB             float64            `bson:"geb,omitempty" json:"geb"`
	GET             float64            `bson:"get,omitempty" json:"get"`
	VET             float64            `bson:"vet,omitempty" json:"vet"`
	Protein         float64            `bson:"protein,omitempty" json:"protein"`
	Lipids          float64            `bson:"lipids,omitempty" json:"lipids"`
	Carbohydrates   float64            `bson:"carbohydrates,omitempty" json:"carbohydrates"`
	CreatedAt       time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
