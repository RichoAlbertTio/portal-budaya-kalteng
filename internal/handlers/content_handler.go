// internal/handlers/content_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"portal-budaya/internal/dto"
	"portal-budaya/internal/models"
	"portal-budaya/internal/util"
)


type ContentHandler struct{ DB *gorm.DB }


func (h *ContentHandler) Create(c *gin.Context) {
var in dto.ContentCreate
if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
content := models.Content{
Title: in.Title,
Slug: util.Slugify(in.Title),
Body: in.Body,
Status: in.Status,
}
if in.ImageURL != nil { content.ImageURL = in.ImageURL }
if in.Summary != nil { content.Summary = in.Summary }
if in.CategoryID != nil { content.CategoryID = in.CategoryID }


// author dari token (opsional) — di demo ini skip
if err := h.DB.Create(&content).Error; err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
// attach tribes & regions bila ada
if len(in.TribeIDs) > 0 { var tribes []models.Tribe; h.DB.Where("id IN ?", in.TribeIDs).Find(&tribes); h.DB.Model(&content).Association("Tribes").Replace(&tribes) }
if len(in.RegionIDs) > 0 { var regs []models.Region; h.DB.Where("id IN ?", in.RegionIDs).Find(&regs); h.DB.Model(&content).Association("Regions").Replace(&regs) }
c.JSON(http.StatusCreated, content)
}


func (h *ContentHandler) Get(c *gin.Context) {
id := c.Param("id")
var content models.Content
if err := h.DB.Preload("Category").Preload("Tribes").Preload("Regions").First(&content, "id = ? OR slug = ?", id, id).Error; err != nil {
c.JSON(http.StatusNotFound, gin.H{"error":"content not found"}); return
}
c.JSON(http.StatusOK, content)
}


func (h *ContentHandler) List(c *gin.Context) {
var items []models.Content
h.DB.Preload("Category").Order("created_at DESC").Find(&items)
c.JSON(http.StatusOK, items)
}