// internal/middlware/basic.go
package middlware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type BasicConfig struct {
Username string
Password string
}


func BasicAuth(cfg BasicConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, p, ok := c.Request.BasicAuth()
		if !ok || u != cfg.Username || p != cfg.Password {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}