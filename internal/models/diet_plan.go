package models

import "time"

// DietPlan defines a meal plan with optional structure.
type DietPlan struct {
	ID                 string    `json:"id" gorm:"primaryKey"`
	PatientID          string    `json:"patientId" gorm:"index;not null"`
	PlanDate           time.Time `json:"planDate"`
	Description        string    `json:"description"`
	Portions           string    `json:"portions"`
	FoodGroups         string    `json:"foodGroups"`
	MealStructure      []Meal    `json:"mealStructure,omitempty" gorm:"foreignKey:DietPlanID;constraint:OnDelete:CASCADE"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

type Meal struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	DietPlanID    string    `json:"dietPlanId" gorm:"index;not null"`
	MealTime      string    `json:"mealTime"`
	MealDescription string  `json:"mealDescription"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type FoodSubstitution struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	MealID           string    `json:"mealId" gorm:"index;not null"`
	OriginalFood     string    `json:"originalFood"`
	SubstituteFood   string    `json:"substituteFood"`
	HouseholdMeasure string    `json:"householdMeasure"`
	WeightGrams      *float64  `json:"weightGrams,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
