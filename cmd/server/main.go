package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"renutri/internal/models"
)

func main() {
	r := gin.Default()

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