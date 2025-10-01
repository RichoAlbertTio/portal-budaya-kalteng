// internal/handlers/about_handler.go
package handlers

import (
	"net/http"

	"portal-budaya/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AboutHandler struct {
	DB *gorm.DB
}

func (h *AboutHandler) Upsert(c *gin.Context) {
	var in struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var a models.About
	// Check if about page already exists
	if err := h.DB.First(&a).Error; err == nil {
		// Update existing
		a.Title = in.Title
		a.Description = in.Description
		h.DB.Save(&a)
		c.JSON(http.StatusOK, a)
		return
	}

	// Create new
	a = models.About{
		Title:       in.Title,
		Description: in.Description,
	}
	h.DB.Create(&a)
	c.JSON(http.StatusCreated, a)
}

func (h *AboutHandler) Get(c *gin.Context) {
	var a models.About
	if err := h.DB.First(&a).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "about not found"})
		return
	}
	c.JSON(http.StatusOK, a)
}