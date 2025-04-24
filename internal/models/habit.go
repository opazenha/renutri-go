package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Habit represents a patient's habit for MongoDB.
type Habit struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID            primitive.ObjectID `bson:"patientId,omitempty" json:"patientId,omitempty"`
	UsualDiet            string             `bson:"usualDiet,omitempty" json:"usualDiet"`
	DietaryRecall        string             `bson:"dietaryRecall,omitempty" json:"dietaryRecall"`
	FrequencyOfConsumption string            `bson:"frequencyOfConsumption,omitempty" json:"frequencyOfConsumption"`
	DietaryRestrictions  string             `bson:"dietaryRestrictions,omitempty" json:"dietaryRestrictions"`
	Supplementation      string             `bson:"supplementation,omitempty" json:"supplementation"`
	Hydration            float64            `bson:"hydration,omitempty" json:"hydration"`
	AlcoholConsumption   string             `bson:"alcoholConsumption,omitempty" json:"alcoholConsumption"`
	SmokingStatus        bool               `bson:"smokingStatus,omitempty" json:"smokingStatus"`
	SleepPattern         string             `bson:"sleepPattern,omitempty" json:"sleepPattern"`
	PhysicalActivity     string             `bson:"physicalActivity,omitempty" json:"physicalActivity"`
	DailyRoutine         string             `bson:"dailyRoutine,omitempty" json:"dailyRoutine"`
	CreatedAt            time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt            time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
