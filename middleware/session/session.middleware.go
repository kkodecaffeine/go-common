package session

import (
	"net/http"
	"time"

	"github.com/kkodecaffeine/go-common/errorcode"
)

var sessions = map[string]session{}

type session struct {
	expiry time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func ValidateSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(errorcode.ACCESS_DENIED.HttpStatusCode)
			return
		}
		w.WriteHeader(errorcode.BAD_REQUEST.HttpStatusCode)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(errorcode.ACCESS_DENIED.HttpStatusCode)
		return
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(errorcode.ACCESS_DENIED.HttpStatusCode)
		return
	}
}
