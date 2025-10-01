// internal/handlers/taxonomy_handler.go (tribes/regions)
package handlers

import (
	"net/http"
	"strings"

	"portal-budaya/internal/models"
	"portal-budaya/internal/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaxHandler struct {
	DB *gorm.DB
}

// ========== TRIBES ==========

func (h *TaxHandler) CreateTribe(c *gin.Context) {
	var in struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tribe := models.Tribe{
		Name: in.Name,
		Slug: util.Slugify(in.Name),
	}
	if in.Description != "" {
		tribe.Description = &in.Description
	}

	if err := h.DB.Create(&tribe).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tribe)
}

func (h *TaxHandler) ListTribe(c *gin.Context) {
	var tribes []models.Tribe
	h.DB.Find(&tribes)
	c.JSON(http.StatusOK, tribes)
}

func (h *TaxHandler) GetTribe(c *gin.Context) {
	slug := c.Param("slug")
	var tribe models.Tribe
	
	// Try slug first, then UUID if it looks like one
	if err := h.DB.Where("slug = ?", slug).First(&tribe).Error; err != nil {
		// If not found by slug, try by UUID (only if it looks like a UUID)
		if len(slug) == 36 && strings.Contains(slug, "-") {
			if err2 := h.DB.Where("id = ?", slug).First(&tribe).Error; err2 != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "tribe not found"})
				return
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "tribe not found"})
			return
		}
	}
	
	c.JSON(http.StatusOK, tribe)
}

func (h *TaxHandler) UpdateTribe(c *gin.Context) {
	id := c.Param("id")
	var tribe models.Tribe
	if err := h.DB.Where("id = ?", id).First(&tribe).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tribe not found"})
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
		tribe.Name = in.Name
		tribe.Slug = util.Slugify(in.Name)
	}
	if in.Description != "" {
		tribe.Description = &in.Description
	}

	if err := h.DB.Save(&tribe).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tribe)
}

func (h *TaxHandler) DeleteTribe(c *gin.Context) {
	id := c.Param("id")
	var tribe models.Tribe
	if err := h.DB.Where("id = ?", id).First(&tribe).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tribe not found"})
		return
	}

	if err := h.DB.Delete(&tribe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "tribe deleted successfully"})
}

// ========== REGIONS ==========

func (h *TaxHandler) CreateRegion(c *gin.Context) {
	var in struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	region := models.Region{
		Name: in.Name,
		Slug: util.Slugify(in.Name),
	}
	if in.Description != "" {
		region.Description = &in.Description
	}

	if err := h.DB.Create(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, region)
}

func (h *TaxHandler) ListRegion(c *gin.Context) {
	var regions []models.Region
	h.DB.Find(&regions)
	c.JSON(http.StatusOK, regions)
}

func (h *TaxHandler) GetRegion(c *gin.Context) {
	slug := c.Param("slug")
	var region models.Region
	
	// Try slug first, then UUID if it looks like one
	if err := h.DB.Where("slug = ?", slug).First(&region).Error; err != nil {
		// If not found by slug, try by UUID (only if it looks like a UUID)
		if len(slug) == 36 && strings.Contains(slug, "-") {
			if err2 := h.DB.Where("id = ?", slug).First(&region).Error; err2 != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "region not found"})
				return
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "region not found"})
			return
		}
	}
	
	c.JSON(http.StatusOK, region)
}

func (h *TaxHandler) UpdateRegion(c *gin.Context) {
	id := c.Param("id")
	var region models.Region
	if err := h.DB.Where("id = ?", id).First(&region).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "region not found"})
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
		region.Name = in.Name
		region.Slug = util.Slugify(in.Name)
	}
	if in.Description != "" {
		region.Description = &in.Description
	}

	if err := h.DB.Save(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, region)
}

func (h *TaxHandler) DeleteRegion(c *gin.Context) {
	id := c.Param("id")
	var region models.Region
	if err := h.DB.Where("id = ?", id).First(&region).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "region not found"})
		return
	}

	if err := h.DB.Delete(&region).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "region deleted successfully"})
}