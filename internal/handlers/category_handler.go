// internal/handlers/category_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"portal-budaya/internal/models"
	"portal-budaya/internal/util"
)

type CategoryHandler struct{ DB *gorm.DB }

func (h *CategoryHandler) Create(c *gin.Context) {
	var in struct{ Name, Description string }
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat := models.Category{Name: in.Name, Slug: util.Slugify(in.Name)}
	if in.Description != "" {
		cat.Description = &in.Description
	}
	if err := h.DB.Create(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

func (h *CategoryHandler) List(c *gin.Context) {
	var cats []models.Category
	h.DB.Find(&cats)
	c.JSON(http.StatusOK, cats)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	if err := h.DB.Where("id = ? OR slug = ?", id, id).First(&cat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	if err := h.DB.Where("id = ?", id).First(&cat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	var in struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if in.Name != "" {
		cat.Name = in.Name
		cat.Slug = util.Slugify(in.Name)
	}
	if in.Description != "" {
		cat.Description = &in.Description
	}

	if err := h.DB.Save(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	if err := h.DB.Where("id = ?", id).First(&cat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	if err := h.DB.Delete(&cat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}