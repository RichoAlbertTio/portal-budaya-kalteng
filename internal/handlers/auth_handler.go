package handlers

import (
	"net/http"
	"strings"
	"time"

	"portal-budaya/internal/dto"
	"portal-budaya/internal/middlware"
	"portal-budaya/internal/models"
	"portal-budaya/internal/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type AuthHandler struct {
DB *gorm.DB
JWTSecret []byte
JWTTTL time.Duration
}


func (h *AuthHandler) Register(c *gin.Context) {
var req dto.RegisterRequest
if err := c.ShouldBindJSON(&req); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
}
pw, _ := util.HashPassword(req.Password)
user := models.User{
Username: req.Username,
Email: strings.ToLower(req.Email),
DisplayName: req.DisplayName,
PasswordHash: pw,
}
if err := h.DB.Create(&user).Error; err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
}
c.JSON(http.StatusCreated, gin.H{"id": user.ID, "username": user.Username})
}


func (h *AuthHandler) Login(c *gin.Context) {
var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	q := h.DB.Where("LOWER(username)=LOWER(?) OR LOWER(email)=LOWER(?)", req.UsernameOrEmail, req.UsernameOrEmail).First(&user)
	if q.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if !util.CheckPassword(user.PasswordHash, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	tok, _ := middlware.GenerateToken(h.JWTSecret, user.ID, user.Role, h.JWTTTL)
	c.JSON(http.StatusOK, gin.H{"access_token": tok, "token_type": "Bearer", "expires_in": int(h.JWTTTL.Seconds())})
}