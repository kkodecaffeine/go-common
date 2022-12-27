package session

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kkodecaffeine/go-common/errorcode"
)

var sessionmap = map[string]session{}

type session struct {
	expiry time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func ValidateSession(c *gin.Context) {

	session := sessions.Default(c)
	sessionID := session.Get("session_token")

	if sessionID == nil {
		c.JSON(http.StatusUnauthorized, nil)

		userSession, exists := sessionmap[sessionID.(string)]
		if !exists {
			c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, nil)
			return
		}

		if userSession.isExpired() {
			delete(sessionmap, sessionID.(string))
			c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, nil)
			return
		}
	}
}
