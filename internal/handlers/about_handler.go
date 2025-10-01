// internal/handlers/about_handler.go
package handlers


import (
"net/http"
"github.com/gin-gonic/gin"
"gorm.io/gorm"
"portal-budaya-kalteng/internal/models"
)


type AboutHandler struct{ DB *gorm.DB }


func (h *AboutHandler) Upsert(c *gin.Context){ var in struct{ Title, Description string }
if err := c.ShouldBindJSON(&in); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
var a models.About
if err := h.DB.First(&a).Error; err==nil {
a.Title=in.Title; a.Description=in.Description; h.DB.Save(&a); c.JSON(http.StatusOK, a); return
}
a = models.About{ Title: in.Title, Description: in.Description }
h.DB.Create(&a); c.JSON(http.StatusCreated, a)
}


func (h *AboutHandler) Get(c *gin.Context){ var a models.About; if err := h.DB.First(&a).Error; err!=nil { c.JSON(http.StatusNotFound, gin.H{"error":"about not found"}); return }; c.JSON(http.StatusOK, a) }