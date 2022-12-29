package token

import (
	"github.com/gin-gonic/gin"
	"github.com/kkodecaffeine/go-common/errorcode"
	"github.com/kkodecaffeine/go-common/rest"
	"github.com/kkodecaffeine/go-common/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidToken(c)
		if err != nil {
			response := rest.NewApiResponse
			response.Error(&errorcode.ACCESS_DENIED, "unauthorized", nil)
			c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, response)
			c.Abort()
			return
		}
		c.Next()
	}
}
