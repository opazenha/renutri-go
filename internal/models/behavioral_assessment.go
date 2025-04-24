package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DayOfWeek represents a day of the week.
type DayOfWeek string

const (
	Monday    DayOfWeek = "segunda-feira"
	Tuesday   DayOfWeek = "terça-feira"
	Wednesday DayOfWeek = "quarta-feira"
	Thursday  DayOfWeek = "quinta-feira"
	Friday    DayOfWeek = "sexta-feira"
	Saturday  DayOfWeek = "sábado"
	Sunday    DayOfWeek = "domingo"
)

// SmokingEntry details the type and frequency of smoking.
type SmokingEntry struct {
	Type            string `bson:"type,omitempty" json:"type,omitempty"`                       // Tipo (e.g., Cigarro, Charuto)
	Frequency       string `bson:"frequency,omitempty" json:"frequency,omitempty"`             // Frequência (e.g., Diária, 3x/semana)
	QuantityPerTime string `bson:"quantityPerTime,omitempty" json:"quantityPerTime,omitempty"` // Quantidade/vez (e.g., 1 maço, 2 unidades)
}

// AlcoholEntry details consumption for a specific type of alcoholic beverage.
type AlcoholEntry struct {
	Type              string   `bson:"type,omitempty" json:"type,omitempty"`                           // Tipo (e.g., Cerveja, Vinho)
	WeeklyFrequency   string   `bson:"weeklyFrequency,omitempty" json:"weeklyFrequency,omitempty"`       // Frequência semanal (e.g., 2x, Fim de semana)
	AmountPerOccasion string   `bson:"amountPerOccasion,omitempty" json:"amountPerOccasion,omitempty"`   // Ingestão por vez (e.g., 3 latas, 2 doses)
	WeeklyVolumeML    *float64 `bson:"weeklyVolumeML,omitempty" json:"weeklyVolumeML,omitempty"`       // Volume ingerido por semana (ml) - *Potentially Calculated*
	WeeklyKcal        *float64 `bson:"weeklyKcal,omitempty" json:"weeklyKcal,omitempty"`               // Kcal provenientes de álcool por semana - *Potentially Calculated*
}

// DailyActivity details physical activity performed on a specific day.
type DailyActivity struct {
	DayOfWeek         DayOfWeek `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`               // Dia da semana (Enum)
	Exercise          string    `bson:"exercise,omitempty" json:"exercise,omitempty"`               // Exercício
	// WeeklyFrequency string    `bson:"weeklyFrequency,omitempty" json:"weeklyFrequency,omitempty"` // Frequência semanal (Seems redundant if listed daily, but included if needed)
	StartTime         string    `bson:"startTime,omitempty" json:"startTime,omitempty"`             // Horário inicial (e.g., "18:00")
	EndTime           string    `bson:"endTime,omitempty" json:"endTime,omitempty"`                 // Horário final (e.g., "19:00")
	Duration          string    `bson:"duration,omitempty" json:"duration,omitempty"`               // Duração (e.g., "1 hora", "45 min") - *Potentially Calculated*
}

// BehavioralAssessment holds lifestyle habit information.
type BehavioralAssessment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`               // MongoDB document ID
	PatientID primitive.ObjectID `bson:"patientId,omitempty" json:"patientId,omitempty"`  // Link to the Patient document
	Date      time.Time          `bson:"date,omitempty" json:"date,omitempty"`            // Date the assessment was performed

	// --- 4.1. Fumo e álcool (Smoking and Alcohol) ---
	DoesSmoke               *bool          `bson:"doesSmoke,omitempty" json:"doesSmoke,omitempty"`                           // Fuma?
	SmokingDetails          []SmokingEntry `bson:"smokingDetails,omitempty" json:"smokingDetails,omitempty"`               // Details if DoesSmoke is true
	ConsumesAlcohol         *bool          `bson:"consumesAlcohol,omitempty" json:"consumesAlcohol,omitempty"`               // Ingere bebidas alcoólicas?
	AlcoholConsumption      []AlcoholEntry `bson:"alcoholConsumption,omitempty" json:"alcoholConsumption,omitempty"`       // Details if ConsumesAlcohol is true
	// TotalWeeklyVolumeML *float64       `bson:"totalWeeklyVolumeML,omitempty" json:"totalWeeklyVolumeML,omitempty"` // *Calculated Field*
	// TotalWeeklyKcal     *float64       `bson:"totalWeeklyKcal,omitempty" json:"totalWeeklyKcal,omitempty"`         // *Calculated Field*

	// --- 4.2. Atividade física (Physical Activity) ---
	PracticesPhysicalActivity *bool           `bson:"practicesPhysicalActivity,omitempty" json:"practicesPhysicalActivity,omitempty"` // Pratica exercícios físicos?
	WeeklyActivitySchedule    []DailyActivity `bson:"weeklyActivitySchedule,omitempty" json:"weeklyActivitySchedule,omitempty"`     // Details if PracticesPhysicalActivity is true

	// Timestamps
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}