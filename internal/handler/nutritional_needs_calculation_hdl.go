package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"renutri/internal/models"
	"renutri/internal/service"
)

type NutritionalNeedsCalculationHandler struct {
	Service *service.NutritionalNeedsCalculationService
}

func NewNutritionalNeedsCalculationHandler(svc *service.NutritionalNeedsCalculationService) *NutritionalNeedsCalculationHandler {
	return &NutritionalNeedsCalculationHandler{Service: svc}
}

func (h *NutritionalNeedsCalculationHandler) Create(c *gin.Context) {
	var input models.NutritionalNeedsCalculation
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

func (h *NutritionalNeedsCalculationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *NutritionalNeedsCalculationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input models.NutritionalNeedsCalculation
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

func (h *NutritionalNeedsCalculationHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *NutritionalNeedsCalculationHandler) List(c *gin.Context) {
	items, err := h.Service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
