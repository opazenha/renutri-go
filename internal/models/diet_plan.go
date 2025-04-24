package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DietPlan defines a meal plan with optional structure.
type DietPlan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID    primitive.ObjectID `bson:"patientId,omitempty" json:"patientId"`
	PlanDate     time.Time          `bson:"planDate,omitempty" json:"planDate"`
	Description  string             `bson:"description,omitempty" json:"description"`
	Portions     string             `bson:"portions,omitempty" json:"portions"`
	FoodGroups   string             `bson:"foodGroups,omitempty" json:"foodGroups"`
	MealStructure []Meal            `bson:"mealStructure,omitempty" json:"mealStructure"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
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
