package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"renutri/internal/db"
	"renutri/internal/handler"
	"renutri/internal/models"
	"renutri/internal/repository"
	"renutri/internal/service"
)

func main() {
	// Connect to MongoDB
	mongoClient, err := db.GetMongoClient()
	if err != nil {
		panic(err)
	}
	database := mongoClient.Database("renutri")

	// Initialize gin router
	r := gin.Default()

	// Initialize repositories/services with the MongoDB database handle
	patientCollection := database.Collection("patients")
	patientRepo := repository.NewPatientRepository(patientCollection)
	patientSvc := service.NewPatientService(patientRepo)
	handler.NewPatientHandler(r, patientSvc)

	appointmentCollection := database.Collection("appointments")
	appointmentRepo := repository.NewAppointmentRepository(appointmentCollection)
	appointmentSvc := service.NewAppointmentService(appointmentRepo)
	handler.NewAppointmentHandler(r, appointmentSvc)

	habitCollection := database.Collection("habits")
	habitRepo := repository.NewHabitRepository(habitCollection)
	habitSvc := service.NewHabitService(habitRepo)
	handler.RegisterHabitRoutes(r, habitSvc)

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