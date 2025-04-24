package models

import "time"

// NutritionalNeedsCalculation records metabolic and nutrient requirements.
type NutritionalNeedsCalculation struct {
	ID                  string    `json:"id" gorm:"primaryKey"`
	PatientID           string    `json:"patientId" gorm:"index;not null"`
	CalculationDate     time.Time `json:"calculationDate"`
	TMB                 float64   `json:"tmb"`
	GEB                 float64   `json:"geb"`
	GET                 float64   `json:"get"`
	VET                 float64   `json:"vet"`
	Carbohydrates       float64   `json:"carbohydrates"`
	Protein             float64   `json:"protein"`
	Lipids              float64   `json:"lipids"`
	MicronutrientNeeds  string    `json:"micronutrientNeeds"`
	WaterIntakeML       float64   `json:"waterIntakeMl"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
