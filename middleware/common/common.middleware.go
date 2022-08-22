package common

import (
	"github.com/gin-gonic/gin"
	"github.com/kkodecaffeine/go-common/rest"
)

func RecoveryMiddleware(c *gin.Context) {
	response := rest.NewApiResponse()

	defer func() {
		err := recover()
		if err != nil {
			// Todo : Send to sentry
			response.Error(&rest.FAILED_INTERNAL_ERROR, "", nil)
			c.JSON(rest.FAILED_INTERNAL_ERROR.HttpStatusCode, response)
			c.Abort()
			return
		}

	}()

	c.Next()
}
