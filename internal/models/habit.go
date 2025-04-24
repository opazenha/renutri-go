package models

import "time"

// Habit holds lifestyle and dietary habits for a patient.
type Habit struct {
	ID                   string    `json:"id" gorm:"primaryKey"`
	PatientID            string    `json:"patientId" gorm:"index;not null"`
	UsualDiet            string    `json:"usualDiet"`
	DietaryRecall        string    `json:"dietaryRecall"`
	FrequencyOfConsumption string  `json:"frequencyOfConsumption"`
	DietaryRestrictions  string    `json:"dietaryRestrictions"`
	Supplementation      string    `json:"supplementation"`
	Hydration            float64   `json:"hydration"`
	AlcoholConsumption   string    `json:"alcoholConsumption"`
	SmokingStatus        bool      `json:"smokingStatus"`
	SleepPattern         string    `json:"sleepPattern"`
	PhysicalActivity     string    `json:"physicalActivity"`
	DailyRoutine         string    `json:"dailyRoutine"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
