package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"renutri/internal/models"
	"renutri/internal/service"
)

type ClinicalEvaluationHandler struct {
	Service *service.ClinicalEvaluationService
}

func NewClinicalEvaluationHandler(svc *service.ClinicalEvaluationService) *ClinicalEvaluationHandler {
	return &ClinicalEvaluationHandler{Service: svc}
}

func (h *ClinicalEvaluationHandler) Create(c *gin.Context) {
	var input models.ClinicalEvaluation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func (h *ClinicalEvaluationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ClinicalEvaluationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input models.ClinicalEvaluation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	objID, err := primitive.ObjectIDFromHex(id)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
    return
}
input.ID = objID
	if err := h.Service.Update(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, input)
}

func (h *ClinicalEvaluationHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *ClinicalEvaluationHandler) List(c *gin.Context) {
	items, err := h.Service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
