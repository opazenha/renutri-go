package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FrameSizeClassification represents body frame size.
type FrameSizeClassification string

const (
	FrameSmall  FrameSizeClassification = "pequena"
	FrameMedium FrameSizeClassification = "média"
	FrameLarge  FrameSizeClassification = "grande"
	FrameUnknown FrameSizeClassification = "desconhecida"
)

// IMCClassification represents the classification based on Body Mass Index.
type IMCClassification string

const (
	IMCUnderweightIII   IMCClassification = "magreza grau III" // Using descriptive terms
	IMCUnderweightII    IMCClassification = "magreza grau II"
	IMCUnderweightI     IMCClassification = "magreza grau I"
	IMCNormal           IMCClassification = "normal"       // NO
	IMCOverweight       IMCClassification = "sobrepeso"    // SP
	IMCObesityI         IMCClassification = "obesidade grau I"
	IMCObesityII        IMCClassification = "obesidade grau II"
	IMCObesityIII       IMCClassification = "obesidade grau III"
	IMCUnknown          IMCClassification = "desconhecido"
)

// ComorbidityRiskClassification represents health risks associated with weight.
type ComorbidityRiskClassification string

const (
	RiskLow         ComorbidityRiskClassification = "baixo"        // B
	RiskMedium      ComorbidityRiskClassification = "médio"        // ME (as per sheet text, though unusual)
	RiskIncreased   ComorbidityRiskClassification = "aumentado"    // A
	RiskModerate    ComorbidityRiskClassification = "moderado"     // MO (as per sheet text)
	RiskSevere      ComorbidityRiskClassification = "severo"       // S
	RiskVerySevere  ComorbidityRiskClassification = "muito severo" // MS
	RiskUnknown     ComorbidityRiskClassification = "desconhecido"
)

// WaistRiskClassification represents health risks associated with waist circumference.
type WaistRiskClassification string

const (
	WaistRiskLow           WaistRiskClassification = "baixo"           // B
	WaistRiskIncreased     WaistRiskClassification = "aumentado"       // A
	WaistRiskVeryIncreased WaistRiskClassification = "muito aumentado" // MA
	WaistRiskUnknown       WaistRiskClassification = "desconhecido"
)

// BodyFatClassification represents the classification based on body fat percentage, age, and sex.
type BodyFatClassification string

const (
	BodyFatVeryLow         BodyFatClassification = "muito baixo"       // MT BXO
	BodyFatExcellent       BodyFatClassification = "excelente"         // EXC
	BodyFatVeryGood        BodyFatClassification = "muito bom"         // MT BOM
	BodyFatGood            BodyFatClassification = "bom"
	BodyFatAdequate        BodyFatClassification = "adequado"          // ADEQ
	BodyFatModeratelyHigh  BodyFatClassification = "moderadamente alto"// MOD ALTO
	BodyFatHigh            BodyFatClassification = "alto"
	BodyFatVeryHigh        BodyFatClassification = "muito alto"        // MT ALTO
	BodyFatUnknown         BodyFatClassification = "desconhecido"
)

// AnthropometricAssessment holds general anthropometric data and a history of evaluations.
type AnthropometricAssessment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`               // MongoDB document ID
	PatientID primitive.ObjectID `bson:"patientId,omitempty" json:"patientId,omitempty"`  // Link to the Patient document

	// --- General Info ---
	// Note: Name, Sex, PracticesPhysicalActivity might be redundant if already in Patient/BehavioralAssessment
	Goal                   string                  `bson:"goal,omitempty" json:"goal,omitempty"`                                   // Objetivo (Overall goal for anthropometry/nutrition plan)
	UsualWeightKg          *float64                `bson:"usualWeightKg,omitempty" json:"usualWeightKg,omitempty"`                 // Peso usual (Kg)
	DesiredWeightKg        *float64                `bson:"desiredWeightKg,omitempty" json:"desiredWeightKg,omitempty"`             // Peso desejado (Kg)
	HeightM                *float64                `bson:"heightM,omitempty" json:"heightM,omitempty"`                             // Estatura (m) - Usually measured once or infrequently
	WristCircumferenceCm   *float64                `bson:"wristCircumferenceCm,omitempty" json:"wristCircumferenceCm,omitempty"`   // Punho (cm)
	FrameSize              FrameSizeClassification `bson:"frameSize,omitempty" json:"frameSize,omitempty"`                         // Compleição ou ossatura - *Calculated: Needs formula using HeightM, WristCircumferenceCm*
	WristDiameterCm        *float64                `bson:"wristDiameterCm,omitempty" json:"wristDiameterCm,omitempty"`             // Diâmetro do punho (cm) - Used in some body composition formulas
	KneeDiameterCm         *float64                `bson:"kneeDiameterCm,omitempty" json:"kneeDiameterCm,omitempty"`               // Diâmetro do joelho (cm) - Used in some body composition formulas

	// --- Evaluation History ---
	Evaluations []EvaluationEntry `bson:"evaluations,omitempty" json:"evaluations,omitempty"` // Slice holding each evaluation snapshot

	// Timestamps
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

// EvaluationEntry holds all measurements and calculations for a single assessment date.
type EvaluationEntry struct {
	EvaluationDate  time.Time `bson:"evaluationDate,omitempty" json:"evaluationDate,omitempty"` // Data da avaliação
	EvaluationTime  string    `bson:"evaluationTime,omitempty" json:"evaluationTime,omitempty"` // Horário (e.g., "10:30")
	AgeAtEvaluation *int      `bson:"ageAtEvaluation,omitempty" json:"ageAtEvaluation,omitempty"` // Idade (Can be calculated: EvaluationDate - Patient.Birthdate)
	IsMenstruating  *bool     `bson:"isMenstruating,omitempty" json:"isMenstruating,omitempty"` // Menstruada (Relevant for females)

	// --- Weight x Height ---
	CurrentWeightKg        *float64                    `bson:"currentWeightKg,omitempty" json:"currentWeightKg,omitempty"`               // Peso atual (Kg)
	IMC                    *float64                    `bson:"imc,omitempty" json:"imc,omitempty"`                                       // IMC (Kg/m²) - *Calculated: CurrentWeightKg / (HeightM * HeightM)*
	TheoreticalMinWeightKg *float64                    `bson:"theoreticalMinWeightKg,omitempty" json:"theoreticalMinWeightKg,omitempty"` // Peso teórico mínimo (Kg) - *Calculated: (18.5 or 22 for elderly) * HeightM^2*
	TheoreticalMaxWeightKg *float64                    `bson:"theoreticalMaxWeightKg,omitempty" json:"theoreticalMaxWeightKg,omitempty"` // Peso teórico máximo (Kg) - *Calculated: (24.9 or 27 for elderly) * HeightM^2*
	IMCClassification      IMCClassification           `bson:"imcClassification,omitempty" json:"imcClassification,omitempty"`           // Classificação do IMC - *Calculated based on IMC, potentially Age*
	ComorbidityRisk        ComorbidityRiskClassification `bson:"comorbidityRisk,omitempty" json:"comorbidityRisk,omitempty"`               // Risco de comorbidades - *Calculated based on IMCClassification*

	// --- Perimeters (cm) ---
	ChestCm                 *float64                `bson:"chestCm,omitempty" json:"chestCm,omitempty"`                                 // Tórax
	ArmRelaxedCm            *float64                `bson:"armRelaxedCm,omitempty" json:"armRelaxedCm,omitempty"`                       // Braço relaxado
	ArmFlexedCm             *float64                `bson:"armFlexedCm,omitempty" json:"armFlexedCm,omitempty"`                         // Braço fletido
	WaistCm                 *float64                `bson:"waistCm,omitempty" json:"waistCm,omitempty"`                                 // Cintura
	AbdomenCm               *float64                `bson:"abdomenCm,omitempty" json:"abdomenCm,omitempty"`                             // Abdômen
	HipCm                   *float64                `bson:"hipCm,omitempty" json:"hipCm,omitempty"`                                     // Quadril
	ThighProximalCm         *float64                `bson:"thighProximalCm,omitempty" json:"thighProximalCm,omitempty"`                 // Coxa proximal
	ThighMedialCm           *float64                `bson:"thighMedialCm,omitempty" json:"thighMedialCm,omitempty"`                     // Coxa medial
	CalfCm                  *float64                `bson:"calfCm,omitempty" json:"calfCm,omitempty"`                                   // Panturrilha
	WaistRiskClassification WaistRiskClassification `bson:"waistRiskClassification,omitempty" json:"waistRiskClassification,omitempty"` // Classificação da cintura - *Calculated based on WaistCm, Sex*

	// --- Skinfolds (mm) ---
	TricepsMm     *float64 `bson:"tricepsMm,omitempty" json:"tricepsMm,omitempty"`         // Triciptal
	PectoralMm    *float64 `bson:"pectoralMm,omitempty" json:"pectoralMm,omitempty"`       // Peitoral
	MidaxillaryMm *float64 `bson:"midaxillaryMm,omitempty" json:"midaxillaryMm,omitempty"` // Axilar média
	SubscapularMm *float64 `bson:"subscapularMm,omitempty" json:"subscapularMm,omitempty"` // Subescapular
	AbdominalMm   *float64 `bson:"abdominalMm,omitempty" json:"abdominalMm,omitempty"`     // Abdominal
	SuprailiacMm  *float64 `bson:"suprailiacMm,omitempty" json:"suprailiacMm,omitempty"`   // Suprailíaca
	ThighMm       *float64 `bson:"thighMm,omitempty" json:"thighMm,omitempty"`             // Coxa
	SkinfoldSum   *float64 `bson:"skinfoldSum,omitempty" json:"skinfoldSum,omitempty"`     // Somatório - *Calculated: Sum of relevant skinfolds based on chosen formula*

	// --- Body Composition ---
	BodyFatPercentage         *float64              `bson:"bodyFatPercentage,omitempty" json:"bodyFatPercentage,omitempty"`                 // % Gordura corporal - *Calculated: Needs specific formula (e.g., Jackson & Pollock, Durnin) using Skinfolds, Age, Sex*
	LeanMassPercentage        *float64              `bson:"leanMassPercentage,omitempty" json:"leanMassPercentage,omitempty"`               // % Peso magro - *Calculated: 100 - BodyFatPercentage*
	FatMassKg                 *float64              `bson:"fatMassKg,omitempty" json:"fatMassKg,omitempty"`                                 // Peso gordo (Kg) - *Calculated: CurrentWeightKg * (BodyFatPercentage / 100)*
	LeanMassKg                *float64              `bson:"leanMassKg,omitempty" json:"leanMassKg,omitempty"`                               // Peso magro (Kg) - *Calculated: CurrentWeightKg - FatMassKg*
	BoneMassKg                *float64              `bson:"boneMassKg,omitempty" json:"boneMassKg,omitempty"`                               // Peso ósseo (Kg) - *Calculated: Needs specific formula (e.g., using Height, Diameters)*
	ResidualMassKg            *float64              `bson:"residualMassKg,omitempty" json:"residualMassKg,omitempty"`                       // Peso residual (Kg) - *Calculated: Needs specific formula (e.g., % of Weight or Weight - Fat - Bone - Muscle)*
	BodyFatClassification     BodyFatClassification `bson:"bodyFatClassification,omitempty" json:"bodyFatClassification,omitempty"`         // Classificação % gordura - *Calculated: Based on BodyFatPercentage, Age, Sex using reference tables (e.g., Pollock 1993)*
	DesiredBodyFatPercentage  *float64              `bson:"desiredBodyFatPercentage,omitempty" json:"desiredBodyFatPercentage,omitempty"`   // % Gord. corporal desejado (Input by nutritionist)
	WeightForDesiredBodyFatKg *float64              `bson:"weightForDesiredBodyFatKg,omitempty" json:"weightForDesiredBodyFatKg,omitempty"` // Peso p/ % gord. desejado - *Calculated: LeanMassKg / (1 - (DesiredBodyFatPercentage / 100))*
}

