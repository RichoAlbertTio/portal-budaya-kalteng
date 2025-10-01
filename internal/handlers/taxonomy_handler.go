// internal/handlers/taxonomy_handler.go (tribes/regions ringkas)
package handlers


import (
"net/http"
"github.com/gin-gonic/gin"
"gorm.io/gorm"
"portal-budaya-kalteng/internal/models"
"portal-budaya-kalteng/internal/util"
)


type TaxHandler struct{ DB *gorm.DB }


func (h *TaxHandler) CreateTribe(c *gin.Context){ var in struct{ Name, Description string }
if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
t:=models.Tribe{Name: in.Name, Slug: util.Slugify(in.Name)}; if in.Description!=""{ t.Description=&in.Description }
if err := h.DB.Create(&t).Error; err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
c.JSON(http.StatusCreated, t) }


func (h *TaxHandler) ListTribe(c *gin.Context){ var out []models.Tribe; h.DB.Find(&out); c.JSON(http.StatusOK, out) }


func (h *TaxHandler) CreateRegion(c *gin.Context){ var in struct{ Name, Description string }
if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
r:=models.Region{Name: in.Name, Slug: util.Slugify(in.Name)}; if in.Description!=""{ r.Description=&in.Description }
if err := h.DB.Create(&r).Error; err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
c.JSON(http.StatusCreated, r) }


func (h *TaxHandler) ListRegion(c *gin.Context){ var out []models.Region; h.DB.Find(&out); c.JSON(http.StatusOK, out) }