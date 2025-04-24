package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AnthropometricAssessment holds body measurements and metrics.
type AnthropometricAssessment struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID    primitive.ObjectID `bson:"patientId,omitempty" json:"patientId"`
	AssessmentDate time.Time        `bson:"assessmentDate,omitempty" json:"assessmentDate"`
	// Add other fields as needed, all with bson/json tags and Mongo-friendly types
	Weight                  float64   `json:"weight"`
	Height                  float64   `json:"height"`
	IdealWeight             float64   `json:"idealWeight"`
	BMI                     float64   `json:"bmi"`
	BMIClassification       string    `json:"bmiClassification"`
	Waist                   *float64  `json:"waist,omitempty"`
	Hip                     *float64  `json:"hip,omitempty"`
	Arm                     *float64  `json:"arm,omitempty"`
	Triceps                 *float64  `json:"triceps,omitempty"`
	Biceps                  *float64  `json:"biceps,omitempty"`
	Subscapular             *float64  `json:"subscapular,omitempty"`
	Suprailiac              *float64  `json:"suprailiac,omitempty"`
	BodyFatPercentage       float64   `json:"bodyFatPercentage"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}
