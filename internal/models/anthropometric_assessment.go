package models

import "time"

// AnthropometricAssessment holds body measurements and metrics.
type AnthropometricAssessment struct {
	ID                      string    `json:"id" gorm:"primaryKey"`
	PatientID               string    `json:"patientId" gorm:"index;not null"`
	AssessmentDate          time.Time `json:"assessmentDate"`
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
