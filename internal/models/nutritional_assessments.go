package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AppetiteClassification describes the patient's appetite.
type AppetiteClassification string

const (
	AppetiteGood      AppetiteClassification = "bom"
	AppetiteFair      AppetiteClassification = "regular"
	AppetitePoor      AppetiteClassification = "ruim"
	AppetiteIncreased AppetiteClassification = "aumentado"
	AppetiteDecreased AppetiteClassification = "diminuído"
	AppetiteUnknown   AppetiteClassification = "desconhecido"
)

// ConsumptionAmount generally describes consumption levels.
type ConsumptionAmount string

const (
	AmountLow      ConsumptionAmount = "baixo" // ou "pouco/a"
	AmountModerate ConsumptionAmount = "moderado/a"
	AmountHigh     ConsumptionAmount = "alto/a" // ou "muito/a"
	AmountUnknown  ConsumptionAmount = "desconhecido"
)

// FoodConsumptionFrequency defines how often a food group is consumed.
type FoodConsumptionFrequency string

const (
	FreqMoreThanOnceDaily FoodConsumptionFrequency = ">1x/dia"
	FreqOnceDaily         FoodConsumptionFrequency = "1x/dia"
	Freq3To6Weekly        FoodConsumptionFrequency = "3 a 6x/semana"
	Freq1To2Weekly        FoodConsumptionFrequency = "1 a 2x/semana"
	FreqBiweekly          FoodConsumptionFrequency = "quinzenal"
	FreqMonthly           FoodConsumptionFrequency = "mensal"
	FreqRarelyNever       FoodConsumptionFrequency = "raro ou nunca"
	FreqUnknown           FoodConsumptionFrequency = "desconhecido"
)

// MealRecordEntry represents a single meal entry in the food record.
type MealRecordEntry struct {
	MealName   string `bson:"mealName,omitempty" json:"mealName,omitempty"`     // Refeições (e.g., Café da manhã, Almoço)
	Time       string `bson:"time,omitempty" json:"time,omitempty"`             // Hora (e.g., "08:00")
	FoodItems  string `bson:"foodItems,omitempty" json:"foodItems,omitempty"`   // Alimentos (Could be string or []string)
	Quantity   string `bson:"quantity,omitempty" json:"quantity,omitempty"`     // Quantidade (e.g., "1 xícara", "100g")
}

// FoodFrequencyEntry represents the consumption frequency of a specific food group.
type FoodFrequencyEntry struct {
	FoodGroup string                   `bson:"foodGroup,omitempty" json:"foodGroup,omitempty"` // Grupo de Alimentos
	Frequency FoodConsumptionFrequency `bson:"frequency,omitempty" json:"frequency,omitempty"` // Frequency enum
}

// NutritionalAssessment holds the detailed food assessment data.
type NutritionalAssessment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`               // MongoDB document ID
	PatientID primitive.ObjectID `bson:"patientId,omitempty" json:"patientId,omitempty"`  // Link to the Patient document
	Date      time.Time          `bson:"date,omitempty" json:"date,omitempty"`            // Date the assessment was performed

	// --- 3. Avaliação Alimentar (Food Assessment) ---
	HadPriorNutritionConsultation *bool                  `bson:"hadPriorNutritionConsultation,omitempty" json:"hadPriorNutritionConsultation,omitempty"` // Já realizou acompanhamento nutricional
	PriorNutritionistName         string                 `bson:"priorNutritionistName,omitempty" json:"priorNutritionistName,omitempty"`             // Com quem
	PriorConsultationGoal         string                 `bson:"priorConsultationGoal,omitempty" json:"priorConsultationGoal,omitempty"`             // Com qual objetivo
	FoodAllergies                 []string               `bson:"foodAllergies,omitempty" json:"foodAllergies,omitempty"`                         // Alergias alimentares
	FoodIntolerances              []string               `bson:"foodIntolerances,omitempty" json:"foodIntolerances,omitempty"`                     // Intolerâncias alimentares
	FoodPreferences               []string               `bson:"foodPreferences,omitempty" json:"foodPreferences,omitempty"`                       // Preferências alimentares
	DislikedFoods                 []string               `bson:"dislikedFoods,omitempty" json:"dislikedFoods,omitempty"`                           // Alimentos que não gosta
	AppetiteClassification        AppetiteClassification `bson:"appetiteClassification,omitempty" json:"appetiteClassification,omitempty"`           // Classificação do apetite (Enum)
	HungriestTime                 string                 `bson:"hungriestTime,omitempty" json:"hungriestTime,omitempty"`                           // Qual horário sente mais fome
	SnacksBetweenMeals            *bool                  `bson:"snacksBetweenMeals,omitempty" json:"snacksBetweenMeals,omitempty"`                   // Belisca entre as refeições
	RecentAppetiteChange          *bool                  `bson:"recentAppetiteChange,omitempty" json:"recentAppetiteChange,omitempty"`             // Alteração recente de apetite
	AppetiteChangeReason          string                 `bson:"appetiteChangeReason,omitempty" json:"appetiteChangeReason,omitempty"`             // Motivo
	WaterIntakePureLiters         *float64               `bson:"waterIntakePureLiters,omitempty" json:"waterIntakePureLiters,omitempty"`           // Quantidade de água pura ingerida por dia (L) - Use pointer for optionality
	TotalFluidIntakeLiters        *float64               `bson:"totalFluidIntakeLiters,omitempty" json:"totalFluidIntakeLiters,omitempty"`         // Quantidade de líquidos ingerida por dia (L) - Use pointer
	DrinksWithMeals               *bool                  `bson:"drinksWithMeals,omitempty" json:"drinksWithMeals,omitempty"`                       // Ingere líquidos junto às refeições
	DrinksWithMealsAmount         string                 `bson:"drinksWithMealsAmount,omitempty" json:"drinksWithMealsAmount,omitempty"`           // Quantidade (e.g., "1 copo", "500ml")
	UsesSupplements               *bool                  `bson:"usesSupplements,omitempty" json:"usesSupplements,omitempty"`                       // Usa suplementos
	SupplementsDetails            string                 `bson:"supplementsDetails,omitempty" json:"supplementsDetails,omitempty"`                 // Quais/indicações de uso
	HouseholdMealPeopleCount      *int                   `bson:"householdMealPeopleCount,omitempty" json:"householdMealPeopleCount,omitempty"`     // Número de pessoas que fazem as refeições na casa - Use pointer
	SaltAmount                    ConsumptionAmount      `bson:"saltAmount,omitempty" json:"saltAmount,omitempty"`                                 // Quantidade de sal (Enum)
	FatAmount                     ConsumptionAmount      `bson:"fatAmount,omitempty" json:"fatAmount,omitempty"`                                   // Quantidade de gordura (Enum)
	MonthlySaltTotalGrams         *float64               `bson:"monthlySaltTotalGrams,omitempty" json:"monthlySaltTotalGrams,omitempty"`           // Total de sal utilizada/mês (g) - Use pointer
	// DailySaltPerPersonGrams    *float64               `bson:"dailySaltPerPersonGrams,omitempty" json:"dailySaltPerPersonGrams,omitempty"`       // *Calculated Field - Omit or store if needed*
	MonthlyOilTotalMilliliters    *float64               `bson:"monthlyOilTotalMilliliters,omitempty" json:"monthlyOilTotalMilliliters,omitempty"` // Total de óleo utilizado/mês (ml) - Use pointer
	// DailyOilPerPersonMilliliters *float64               `bson:"dailyOilPerPersonMilliliters,omitempty" json:"dailyOilPerPersonMilliliters,omitempty"` // *Calculated Field - Omit or store if needed*
	SaltFatObservations           string                 `bson:"saltFatObservations,omitempty" json:"saltFatObservations,omitempty"`               // Obs.:
	// SaltIntakeClassification   ConsumptionAmount      `bson:"saltIntakeClassification,omitempty" json:"saltIntakeClassification,omitempty"`     // *Likely calculated from DailySaltPerPersonGrams - Omit or store if needed*

	// --- 3.1. Registro Alimentar (Food Record) ---
	FoodRecord []MealRecordEntry `bson:"foodRecord,omitempty" json:"foodRecord,omitempty"` // Slice of meal entries

	// --- 3.2. Frequência Alimentar (Food Frequency) ---
	FoodFrequency             []FoodFrequencyEntry `bson:"foodFrequency,omitempty" json:"foodFrequency,omitempty"`         // Slice of food group frequencies
	FoodFrequencyObservations string               `bson:"foodFrequencyObservations,omitempty" json:"foodFrequencyObservations,omitempty"` // Obs.:

	// Timestamps
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}