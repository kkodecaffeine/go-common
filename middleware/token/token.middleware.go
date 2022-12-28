package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkodecaffeine/go-common/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
