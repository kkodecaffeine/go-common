package session

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kkodecaffeine/go-common/errorcode"
	"github.com/kkodecaffeine/go-common/rest"
)

var sessionmap = map[string]session{}

type session struct {
	expiry time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func ValidateSession(c *gin.Context) {
	response := rest.NewApiResponse()

	session := sessions.Default(c)
	sessionID := session.Get("session_id")

	if sessionID == nil {
		response.Error(&errorcode.ACCESS_DENIED, "", nil)
		c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, response)

		userSession, exists := sessionmap[sessionID.(string)]
		if !exists {
			c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, response)
			return
		}

		if userSession.isExpired() {
			delete(sessionmap, sessionID.(string))
			c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, response)
			return
		}
	}
}
