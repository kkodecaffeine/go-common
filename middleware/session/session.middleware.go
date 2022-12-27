package session

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
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

	var store = sessions.NewCookieStore([]byte("secret"))
	sessionID, _ := store.Get(c.Request, "session")

	if sessionID == nil {
		response.Error(&errorcode.ACCESS_DENIED, "", nil)
		c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, response)
		return
	}

	userSession, exists := sessionmap[sessionID.Name()]
	if !exists {
		c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, response)
		return
	}

	if userSession.isExpired() {
		delete(sessionmap, sessionID.ID)
		c.JSON(errorcode.ACCESS_DENIED.HttpStatusCode, response)
		return
	}
}
