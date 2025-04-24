package handler

import (
	"net/http"
	"renutri/internal/models"
	"renutri/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

type HabitHandler struct {
	Service *service.HabitService
}

func RegisterHabitRoutes(r *gin.Engine, svc *service.HabitService) {
  h := NewHabitHandler(svc)
  grp := r.Group("/habits")
  grp.POST("",    h.Create)
  grp.GET("",     h.List)
  grp.GET("/:id", h.GetByID)
  grp.PUT("/:id", h.Update)
  grp.DELETE("/:id", h.Delete)
}

func NewHabitHandler(svc *service.HabitService) *HabitHandler {
	return &HabitHandler{Service: svc}
}

func (h *HabitHandler) Create(c *gin.Context) {
	var input models.Habit
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

func (h *HabitHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *HabitHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input models.Habit
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

func (h *HabitHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *HabitHandler) List(c *gin.Context) {
	items, err := h.Service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
