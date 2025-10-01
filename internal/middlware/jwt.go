// internal/middleware/jwt.go
package middleware


import (
"net/http"
"time"


"github.com/gin-gonic/gin"
"github.com/golang-jwt/jwt/v5"
)


type JWTConfig struct {
Secret []byte
}


func GenerateToken(secret []byte, userID, role string, ttl time.Duration) (string, error) {
claims := jwt.MapClaims{
"sub": userID,
"role": role,
"exp": time.Now().Add(ttl).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
return token.SignedString(secret)
}


func JWTAuth(cfg JWTConfig) gin.HandlerFunc {
return func(c *gin.Context) {
tokStr := c.GetHeader("Authorization")
if !strings.HasPrefix(tokStr, "Bearer ") {
c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"missing bearer token"})
return
}
tokStr = strings.TrimPrefix(tokStr, "Bearer ")
tok, err := jwt.Parse(tokStr, func(t *jwt.Token) (interface{}, error) { return cfg.Secret, nil })
if err != nil || !tok.Valid {
c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
return
}
c.Set("claims", tok.Claims)
c.Next()
}
}