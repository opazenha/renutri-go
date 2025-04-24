package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BowelFunction describes the patient's typical bowel function.
type BowelFunction string

const (
	BowelNormal      BowelFunction = "normal"
	BowelConstipated BowelFunction = "obstipado"
	BowelDiarrhea    BowelFunction = "diarreia"
	BowelIrregular   BowelFunction = "irregular"
	BowelOther       BowelFunction = "outro"
	BowelDiarrheal   BowelFunction = "diarreico"
	BowelMixed       BowelFunction = "misto"
	BowelUnknown     BowelFunction = "desconhecido"
)

// UrineColor describes the typical urine color.
type UrineColor string

const (
	UrineClear       UrineColor = "clara"
	UrineLightYellow UrineColor = "amarela clara"
	UrineDarkYellow  UrineColor = "amarela escura"
	UrineOrange      UrineColor = "laranja"
	UrineOther       UrineColor = "outra"
	UrineUnknown     UrineColor = "desconhecida"
)

// StressLevel indicates the patient's perceived stress level.
type StressLevel string

const (
	StressNone     StressLevel = "nenhum"
	StressLow      StressLevel = "pouco"
	StressModerate StressLevel = "moderado"
	StressHigh     StressLevel = "alto"
	StressExtreme  StressLevel = "extremo"
	StressUnknown  StressLevel = "desconhecido"
)

// ClinicalEvaluation represents a clinical evaluation for MongoDB.
type ClinicalEvaluation struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID               primitive.ObjectID `bson:"patientId,omitempty" json:"patientId"`
	EvaluationDate          time.Time          `bson:"evaluationDate,omitempty" json:"evaluationDate"`
	Symptoms                string             `bson:"symptoms,omitempty" json:"symptoms"`
	LaboratoryResults       string             `bson:"laboratoryResults,omitempty" json:"laboratoryResults"`
	CreatedAt               time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt               time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`

	// --- Avaliação Clínica (Clinical Evaluation) ---
	MainComplaints          []string `bson:"main_complaints,omitempty" json:"main_complaints,omitempty"`
	ComplaintsDuration      string   `bson:"complaints_duration,omitempty" json:"complaints_duration,omitempty"`
	BowelFunction           BowelFunction   `bson:"bowel_function,omitempty" json:"bowel_function,omitempty"`
	UrineColor              UrineColor   `bson:"urine_color,omitempty" json:"urine_color,omitempty"`
	EvacuationFrequency     string   `bson:"evacuation_frequency,omitempty" json:"evacuation_frequency,omitempty"`
	HasEdema                *bool    `bson:"has_edema,omitempty" json:"has_edema,omitempty"`
	StressLevel             StressLevel	   `bson:"stress_level,omitempty" json:"stress_level,omitempty"`
	SleepHours              string   `bson:"sleep_hours,omitempty" json:"sleep_hours,omitempty"`
	SleepsWell              *bool    `bson:"sleeps_well,omitempty" json:"sleeps_well,omitempty"`
	SleepQualityReason      string   `bson:"sleep_quality_reason,omitempty" json:"sleep_quality_reason,omitempty"`
	EatingDisorders         []string `bson:"eating_disorders,omitempty" json:"eating_disorders,omitempty"`
	HadRecentExams          *bool    `bson:"had_recent_exams,omitempty" json:"had_recent_exams,omitempty"`
	RecentExamsDetails      string   `bson:"recent_exams_details,omitempty" json:"recent_exams_details,omitempty"`
	HadRecentTreatments     *bool    `bson:"had_recent_treatments,omitempty" json:"had_recent_treatments,omitempty"`
	RecentTreatmentsDetails string   `bson:"recent_treatments_details,omitempty" json:"recent_treatments_details,omitempty"`
	HadRecentSurgeries      *bool    `bson:"had_recent_surgeries,omitempty" json:"had_recent_surgeries,omitempty"`
	RecentSurgeriesDetails  string   `bson:"recent_surgeries_details,omitempty" json:"recent_surgeries_details,omitempty"`
	OtherComplaints         string   `bson:"other_complaints,omitempty" json:"other_complaints,omitempty"`
	FamilyDiseases          []string `bson:"family_diseases,omitempty" json:"family_diseases,omitempty"`
	CurrentMedications      []string `bson:"current_medications,omitempty" json:"current_medications,omitempty"`
	SystolicPressure        int      `bson:"systolic_pressure,omitempty" json:"systolic_pressure,omitempty"`
	DiastolicPressure       int      `bson:"diastolic_pressure,omitempty" json:"diastolic_pressure,omitempty"`
}
