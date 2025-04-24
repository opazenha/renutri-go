package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"renutri/internal/config"
	"renutri/internal/db"
	"renutri/internal/handler"
	"renutri/internal/models"
	"renutri/internal/repository"
	"renutri/internal/service"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Connect database
	dbConn, err := db.Open(cfg.DBUrl)
	if err != nil {
		panic(err)
	}

	// clean legacy schema and migrate all relational tables
	if err := dbConn.Exec("DROP TABLE IF EXISTS patients CASCADE").Error; err != nil {
		panic(err)
	}
	if err := dbConn.AutoMigrate(
		&models.Patient{},
		&models.Appointment{},
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
	); err != nil {
		panic(fmt.Errorf("auto-migrate models: %w", err))
	}

	r := gin.Default()

	// Setup repository, service, handler
	repo := repository.NewPatientRepository(dbConn)
	svc := service.NewPatientService(repo)
	handler.NewPatientHandler(r, svc)

	repoH := repository.NewHabitRepository(dbConn)
	svcH := service.NewHabitService(repoH)
	handler.RegisterHabitRoutes(r, svcH)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "System is HEALTHY.",
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		var patient models.Patient
		if err := ctx.BindJSON(&patient); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Patient: ", patient)
		ctx.JSON(http.StatusOK, patient)
	})

	fmt.Println("Server running on port 8881")
	r.Run(":8881")
}