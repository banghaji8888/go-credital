package utils

import (
	"go-credital/models"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

var sessionID = "credital-sessions"

// SetSession - set sessions
func SetSession(c echo.Context, key string, user models.SysUser) {
	sess, _ := session.Get(sessionID, c)
	sess.Values[key] = user
	sess.Save(c.Request(), c.Response())
}

// GetSession - get sessions data
func GetSession(c echo.Context) *sessions.Session {
	sess, err := session.Get(sessionID, c)

	if err != nil {
		return nil
	}

	return sess
}

// DeleteSession - delete current session
func DeleteSession(c echo.Context) {
	session, _ := session.Get(sessionID, c)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
}
