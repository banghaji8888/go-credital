package main

import (
	"go-credital/render"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type M map[string]interface{}

func main() {
	e := echo.New()
	e.Use(middleware.Static("/assets"))
	e.Renderer = render.NewRenderer("views", true)

	e.GET("/", func(c echo.Context) error {
		data := M{"message": "Hello World!"}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.GET("/login", func(c echo.Context) error {
		data := M{"message": "Hello World!"}
		return c.Render(http.StatusOK, "login.html", data)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
