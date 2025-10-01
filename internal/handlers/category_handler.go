// internal/handlers/category_handler.go
package handlers


import (
"net/http"


"github.com/gin-gonic/gin"
"gorm.io/gorm"


"portal-budaya-kalteng/internal/models"
"portal-budaya-kalteng/internal/util"
)


type CategoryHandler struct{ DB *gorm.DB }


func (h *CategoryHandler) Create(c *gin.Context) {
var in struct{ Name, Description string }
if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
cat := models.Category{ Name: in.Name, Slug: util.Slugify(in.Name) }
if in.Description != "" { cat.Description = &in.Description }
if err := h.DB.Create(&cat).Error; err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
c.JSON(http.StatusCreated, cat)
}


func (h *CategoryHandler) List(c *gin.Context) {
var cats []models.Category
h.DB.Find(&cats)
c.JSON(http.StatusOK, cats)
}