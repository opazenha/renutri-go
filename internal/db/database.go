package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"renutri/internal/models"
)

// Open initializes and returns a GORM DB connection.
func Open(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    // Drop all existing tables for a full reset
    if err := db.Migrator().DropTable(
        &models.Appointment{},
        &models.GeneratedDocument{},
        &models.Guidance{},
        &models.FoodSubstitution{},
        &models.Meal{},
        &models.DietPlan{},
        &models.NutritionalNeedsCalculation{},
        &models.AnthropometricAssessment{},
        &models.ClinicalEvaluation{},
        &models.Habit{},
        &models.HealthCondition{},
        &models.Patient{},
    ); err != nil {
        return nil, err
    }
    // AutoMigrate all models
    err = db.AutoMigrate(
        &models.Patient{},
        &models.HealthCondition{},
        &models.Habit{},
        &models.ClinicalEvaluation{},
        &models.AnthropometricAssessment{},
        &models.NutritionalNeedsCalculation{},
        &models.DietPlan{},
        &models.Meal{},
        &models.FoodSubstitution{},
        &models.Guidance{},
        &models.GeneratedDocument{},
        &models.Appointment{},
    )
    if err != nil {
        return nil, err
    }
    return db, nil
}
