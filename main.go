package main

import (
	"fmt"
	"go-credital/render"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

var store = sessions.NewCookieStore([]byte("rahasiaitusecret"))
var SESSION_ID = "credital-sessions"

type M map[string]interface{}

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	e.Renderer = render.NewRenderer("views", true)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/login")
	})

	e.GET("/login", func(c echo.Context) error {
		data := M{}
		return c.Render(http.StatusOK, "login", data)
	})

	e.POST("/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username == "youngky" && password == "andreas" {
			session, _ := store.Get(c.Request(), SESSION_ID)
			session.Values["username"] = username
			session.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusFound, "/dashboard")
		}

		data := M{
			"error":    "Username dan Password tidak sama",
			"username": username,
		}

		return c.Render(http.StatusOK, "login.html", data)
	})

	e.GET("/dashboard", func(c echo.Context) error {
		data := M{"message": "Hello World!"}
		return c.Render(http.StatusOK, "dashboard", data)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := store.Get(c.Request(), SESSION_ID)
			if len(session.Values) == 0 {
				return c.Redirect(http.StatusFound, "/login")
			}

			fmt.Println(session.Values["username"])
			return next(c)
		}
	})

	e.GET("/logout", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusFound, "/login")
	})

	e.Logger.Fatal(e.Start(":9000"))
}

/*package main

import (
  "fmt"
  "reflect"
)

// Name of the struct tag used in examples
const tagName = "validate"

type User struct {
  Id    int    `validate:"-"`
  Name  string `validate:"presence,min=2,max=32"`
  Email string `validate:"email,required"`
}

func main() {
  user := User{
    Id:    1,
    Name:  "John Doe",
    Email: "john@example",
  }

  // TypeOf returns the reflection Type that represents the dynamic type of variable.
  // If variable is a nil interface value, TypeOf returns nil.
  t := reflect.TypeOf(user)

  // Get the type and kind of our user variable
  fmt.Println("Type:", t.Name())
  fmt.Println("Kind:", t.Kind())

  // Iterate over all available fields and read the tag value
  for i := 0; i < t.NumField(); i++ {
    // Get the field, returns https://golang.org/pkg/reflect/#StructField
    field := t.Field(i)

    // Get the field tag value
    tag := field.Tag.Get(tagName)

	fmt.Println(tag)
	output:
	 	Type: User
		Kind: struct
		-
		presence,min=2,max=32
		email,required
  }
}*/
