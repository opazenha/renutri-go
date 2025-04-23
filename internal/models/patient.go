package models

import "time"

// Data structure for a Patient
// Central entity holding all patient information

// Patient represents a nutritionist's patient record
// with personal, medical, dietary, and plan details.
type Patient struct {
    PatientID                      string                         `json:"patientId"`
    PersonalData                   PersonalData                   `json:"personalData"`
    HealthHistory                  []HealthCondition              `json:"healthHistory"`
    Habits                         Habits                         `json:"habits"`
    ClinicalEvaluations            []ClinicalEvaluation           `json:"clinicalEvaluations"`
    AnthropometricAssessments      []AnthropometricAssessment     `json:"anthropometricAssessments"`
    NutritionalNeedsCalculations   []NutritionalNeedsCalculation  `json:"nutritionalNeedsCalculations"`
    DietPlans                      []DietPlan                     `json:"dietPlans"`
    NutritionalGuidance            []Guidance                     `json:"nutritionalGuidance"`
    GeneratedDocuments             []GeneratedDocument            `json:"generatedDocuments"`
}

// PersonalData holds basic demographic and contact info
// for a patient.
type PersonalData struct {
    Name         string    `json:"name"`
    DateOfBirth  time.Time `json:"dateOfBirth"`
    Profession   string    `json:"profession"`
    Address      string    `json:"address"`
    Contact      Contact   `json:"contact"`
}

type Contact struct {
    Phone string `json:"phone"`
    Email string `json:"email"`
}

// HealthCondition records a medical condition
// with optional diagnosis date and medication details.
type HealthCondition struct {
    ConditionName string     `json:"conditionName"`
    DiagnosisDate *time.Time `json:"diagnosisDate,omitempty"`
    Status        string     `json:"status"`
    Medication    string     `json:"medication,omitempty"`
}

// Habits captures lifestyle factors and dietary habits.
type Habits struct {
    DietaryHabits       DietaryHabits `json:"dietaryHabits"`
    Hydration           float64       `json:"hydration"`
    AlcoholConsumption  string        `json:"alcoholConsumption"`
    SmokingStatus       bool          `json:"smokingStatus"`
    SleepPattern        string        `json:"sleepPattern"`
    PhysicalActivity    string        `json:"physicalActivity"`
    DailyRoutine        string        `json:"dailyRoutine"`
}

type DietaryHabits struct {
    UsualDiet              string `json:"usualDiet"`
    DietaryRecall          string `json:"dietaryRecall"`
    FrequencyOfConsumption string `json:"frequencyOfConsumption"`
    DietaryRestrictions    string `json:"dietaryRestrictions"`
    Supplementation        string `json:"supplementation"`
}

// ClinicalEvaluation stores assessment date,
// symptoms, and lab results.
type ClinicalEvaluation struct {
    EvaluationDate   time.Time `json:"evaluationDate"`
    Symptoms         string    `json:"symptoms"`
    LaboratoryResults string   `json:"laboratoryResults"`
}

// AnthropometricAssessment holds body measurements
// and calculated metrics.
type AnthropometricAssessment struct {
    AssessmentDate             time.Time       `json:"assessmentDate"`
    Weight                     float64         `json:"weight"`
    Height                     float64         `json:"height"`
    IdealWeight                float64         `json:"idealWeight"`
    BMI                        float64         `json:"bmi"`
    BMIClassification          string          `json:"bmiClassification"`
    Circumferences             Circumferences  `json:"circumferences"`
    CircumferenceClassification string         `json:"circumferenceClassification"`
    Skinfolds                  Skinfolds       `json:"skinfolds"`
    SkinfoldClassification     string          `json:"skinfoldClassification"`
    BodyFatPercentage          float64         `json:"bodyFatPercentage"`
}

type Circumferences struct {
    Waist *float64 `json:"waist,omitempty"`
    Hip   *float64 `json:"hip,omitempty"`
    Arm   *float64 `json:"arm,omitempty"`
}

type Skinfolds struct {
    Triceps      *float64 `json:"triceps,omitempty"`
    Biceps       *float64 `json:"biceps,omitempty"`
    Subscapular  *float64 `json:"subscapular,omitempty"`
    Suprailiac   *float64 `json:"suprailiac,omitempty"`
}

// NutritionalNeedsCalculation records metabolic
// and macro/micronutrient requirements.
type NutritionalNeedsCalculation struct {
    CalculationDate     time.Time            `json:"calculationDate"`
    TMB                 float64              `json:"tmb"`
    GEB                 float64              `json:"geb"`
    GET                 float64              `json:"get"`
    VET                 float64              `json:"vet"`
    MacronutrientNeeds  MacronutrientNeeds  `json:"macronutrientNeeds"`
    MicronutrientNeeds  string               `json:"micronutrientNeeds"`
    WaterIntakeML       float64              `json:"waterIntakeMl"`
}

type MacronutrientNeeds struct {
    Carbohydrates float64 `json:"carbohydrates"`
    Protein       float64 `json:"protein"`
    Lipids        float64 `json:"lipids"`
}

// DietPlan defines a meal plan with optional structure
// and substitutions.
type DietPlan struct {
    PlanDate            time.Time            `json:"planDate"`
    Description         string               `json:"description"`
    Portions            string               `json:"portions"`
    FoodGroups          string               `json:"foodGroups"`
    MealStructure       []Meal               `json:"mealStructure,omitempty"`
    FoodSubstitutionList []FoodSubstitution `json:"foodSubstitutionList,omitempty"`
}

type Meal struct {
    MealTime          string               `json:"mealTime"`
    MealDescription   string               `json:"mealDescription"`
    FoodSubstitutions []FoodSubstitution   `json:"foodSubstitutions,omitempty"`
}

type FoodSubstitution struct {
    OriginalFood      string   `json:"originalFood"`
    SubstituteFood    string   `json:"substituteFood"`
    HouseholdMeasure  string   `json:"householdMeasure"`
    WeightGrams       *float64 `json:"weightGrams,omitempty"`
}

// Guidance for nutritionist advice or instructions.
type Guidance struct {
    GuidanceDate time.Time `json:"guidanceDate"`
    Content      string    `json:"content"`
}

// GeneratedDocument links to produced files (PDFs, etc.).
type GeneratedDocument struct {
    DocumentDate time.Time `json:"documentDate"`
    DocumentType string    `json:"documentType"`
    FilePath     string    `json:"filePath"`
}