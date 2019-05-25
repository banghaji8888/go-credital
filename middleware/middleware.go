package middleware

import (
	"go-credital/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// CustomServerHeader => edit response server
func CustomServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "CreditalPanel/1.0")
		return next(c)
	}
}

//AuthUser - auth user
func AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := utils.GetSession(c)
		if len(session.Values) == 0 {
			return c.Redirect(http.StatusFound, "/login")
		}

		return next(c)
	}
}

// GeneralError => default response if panic
func GeneralError(err error, c echo.Context) {
	log.Println("Fatal error: bro", err, c.Request().URL)
	c.String(http.StatusInternalServerError, "error gan")
}
