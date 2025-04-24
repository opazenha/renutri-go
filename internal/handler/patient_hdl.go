// PatientHandler registers CRUD routes for patients.
package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "renutri/internal/models"
    "renutri/internal/service"
)

// NewPatientHandler binds patient routes to the Gin engine.
func NewPatientHandler(r *gin.Engine, svc service.PatientService) {
    h := &patientHandler{svc: svc}
    grp := r.Group("/patients")
    grp.POST("", h.Create)
    grp.GET("", h.List)
    grp.GET(":id", h.GetByID)
    grp.PUT(":id", h.Update)
    grp.DELETE(":id", h.Delete)
}

type patientHandler struct {
    svc service.PatientService
}

func (h *patientHandler) Create(c *gin.Context) {
    var p models.Patient
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.svc.CreatePatient(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, p)
}

func (h *patientHandler) List(c *gin.Context) {
    list, err := h.svc.ListPatients()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, list)
}

func (h *patientHandler) GetByID(c *gin.Context) {
    id := c.Param("id")
    p, err := h.svc.GetPatient(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, p)
}

func (h *patientHandler) Update(c *gin.Context) {
    id := c.Param("id")
    var p models.Patient
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    objID, err := primitive.ObjectIDFromHex(id)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
    return
}
p.ID = objID
    if err := h.svc.UpdatePatient(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, p)
}

func (h *patientHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.svc.DeletePatient(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}