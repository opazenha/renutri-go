package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"renutri/internal/models"
	"renutri/internal/service"
)

// NewAppointmentHandler binds appointment routes to the Gin engine.
func NewAppointmentHandler(r *gin.Engine, svc service.AppointmentService) {
	h := &appointmentHandler{svc: svc}
	routes := r.Group("/appointments")
	routes.POST("", h.Create)
	routes.GET("", h.List)
	routes.GET("/:id", h.GetByID)
	routes.PUT("/:id", h.Update)
	routes.DELETE("/:id", h.Delete)

	// List appointments by patient
	r.GET("/patients/:patientId/appointments", h.ListByPatient)
}

type appointmentHandler struct {
	svc service.AppointmentService
}

// Create handles POST /appointments
func (h *appointmentHandler) Create(c *gin.Context) {
	var app models.Appointment
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.CreateAppointment(&app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, app)
}

// List handles GET /appointments
func (h *appointmentHandler) List(c *gin.Context) {
	apps, err := h.svc.ListAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, apps)
}

// ListByPatient handles GET /patients/:patientId/appointments
func (h *appointmentHandler) ListByPatient(c *gin.Context) {
	pid := c.Param("patientId")
	apps, err := h.svc.ListByPatient(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, apps)
}

// GetByID handles GET /appointments/:id
func (h *appointmentHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	app, err := h.svc.GetAppointment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, app)
}

// Update handles PUT /appointments/:id
func (h *appointmentHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var app models.Appointment
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	app.ID = objID
	if err := h.svc.UpdateAppointment(&app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, app)
}

// Delete handles DELETE /appointments/:id
func (h *appointmentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.DeleteAppointment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
