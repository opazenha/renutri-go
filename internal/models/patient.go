package models

import (
	"time"
)

// Data structure for a Patient
// Central entity holding all patient information

// Patient represents a nutritionist's patient record with full relational sub-entities.
type Patient struct {
	ID                 string                   `json:"id" gorm:"primaryKey"`

	// --- Personal Data ---
	Name                string                   `json:"name,omitempty"`
	Birthdate           time.Time                `json:"birthdate,omitempty"`
	Age                 int                      `json:"age,omitempty"`
	Sex                 Gender                   `json:"sex,omitempty"`
	EducationLevel      EducationLevel           `json:"educationLevel,omitempty"`
	MaritalStatus       MaritalStatus            `json:"maritalStatus,omitempty"`
	Profession          string                   `json:"profession,omitempty"`
	Address             string                   `json:"address,omitempty"`
	Phones              []string                 `json:"phones,omitempty" gorm:"type:text[]"`
	Email               string                   `json:"email,omitempty"`
	HealthInsurance     string                   `json:"healthInsurance,omitempty"`
	Referral            string                   `json:"referral,omitempty"`
	ConsultationReason  string                   `json:"consultationReason,omitempty"`
	Observations        string                   `json:"observations,omitempty"`

	// --- Clinical Data Relations ---
	HealthConditions            []HealthCondition           `json:"healthConditions,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	Habit                       Habit                       `json:"habit,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	ClinicalEvaluations         []ClinicalEvaluation        `json:"clinicalEvaluations,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	AnthropometricAssessments   []AnthropometricAssessment  `json:"anthropometricAssessments,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	NutritionalNeedsCalculations []NutritionalNeedsCalculation `json:"nutritionalNeedsCalculations,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	DietPlans                   []DietPlan                  `json:"dietPlans,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	GuidanceRecords            []Guidance                  `json:"guidanceRecords,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	GeneratedDocuments         []GeneratedDocument         `json:"generatedDocuments,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	Appointments                []Appointment               `json:"appointments,omitempty" gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`

	// Timestamps
	CreatedAt         time.Time                `json:"createdAt"`
	UpdatedAt         time.Time                `json:"updatedAt"`
}