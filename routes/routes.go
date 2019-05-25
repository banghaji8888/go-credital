package routes

import (
	"go-credital/controllers"
	m "go-credital/middleware"
	"go-credital/render"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

// Init -->>
func Init() *echo.Echo {
	e := echo.New()

	// static folder
	e.Static("/static", "assets")

	// custom render
	e.Renderer = render.NewRenderer("views", true)

	// middleware
	e.HTTPErrorHandler = m.GeneralError
	e.Use(middleware.CORS())
	e.Use(m.CustomServerHeader)
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secretiturahasia"))))

	// default route
	e.GET("/", func(c echo.Context) error { return c.Redirect(http.StatusFound, "/login") })
	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.DoLogin)
	e.GET("/dashboard", controllers.Dashboard, m.AuthUser)
	e.GET("/logout", controllers.Logout)

	return e
}
