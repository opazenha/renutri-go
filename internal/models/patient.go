package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Patient represents a nutritionist's patient record for MongoDB.
type Patient struct {
	ID                 primitive.ObjectID   `bson:"_id,omitempty" json:"id"`

	// --- Personal Data ---
	Name                string              `bson:"name,omitempty" json:"name,omitempty"`
	Birthdate           time.Time           `bson:"birthdate,omitempty" json:"birthdate,omitempty"`
	Age                 int                 `bson:"age,omitempty" json:"age,omitempty"`
	Sex                 string              `bson:"sex,omitempty" json:"sex,omitempty"`
	EducationLevel      string              `bson:"educationLevel,omitempty" json:"educationLevel,omitempty"`
	MaritalStatus       string              `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	Profession          string              `bson:"profession,omitempty" json:"profession,omitempty"`
	Address             string              `bson:"address,omitempty" json:"address,omitempty"`
	Phones              []string            `bson:"phones,omitempty" json:"phones,omitempty"`
	Email               string              `bson:"email,omitempty" json:"email,omitempty"`
	HealthInsurance     string              `bson:"healthInsurance,omitempty" json:"healthInsurance,omitempty"`
	Referral            string              `bson:"referral,omitempty" json:"referral,omitempty"`
	ConsultationReason  string              `bson:"consultationReason,omitempty" json:"consultationReason,omitempty"`
	Observations        string              `bson:"observations,omitempty" json:"observations,omitempty"`

	// --- Clinical Data Relations ---
	// For MongoDB, these can be []primitive.ObjectID or embedded structs, depending on use-case.
	HealthConditionIDs            []primitive.ObjectID   `bson:"healthConditionIds,omitempty" json:"healthConditionIds,omitempty"`
	HabitID                      *primitive.ObjectID    `bson:"habitId,omitempty" json:"habitId,omitempty"`
	ClinicalEvaluationIDs         []primitive.ObjectID   `bson:"clinicalEvaluationIds,omitempty" json:"clinicalEvaluationIds,omitempty"`
	AnthropometricAssessmentIDs   []primitive.ObjectID   `bson:"anthropometricAssessmentIds,omitempty" json:"anthropometricAssessmentIds,omitempty"`
	NutritionalNeedsCalculationIDs []primitive.ObjectID  `bson:"nutritionalNeedsCalculationIds,omitempty" json:"nutritionalNeedsCalculationIds,omitempty"`
	DietPlanIDs                   []primitive.ObjectID   `bson:"dietPlanIds,omitempty" json:"dietPlanIds,omitempty"`
	GuidanceRecordIDs             []primitive.ObjectID   `bson:"guidanceRecordIds,omitempty" json:"guidanceRecordIds,omitempty"`
	GeneratedDocumentIDs          []primitive.ObjectID   `bson:"generatedDocumentIds,omitempty" json:"generatedDocumentIds,omitempty"`
	AppointmentIDs                []primitive.ObjectID   `bson:"appointmentIds,omitempty" json:"appointmentIds,omitempty"`

	// Timestamps
	CreatedAt         time.Time             `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt         time.Time             `bson:"updatedAt,omitempty" json:"updatedAt"`
}