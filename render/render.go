package render

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *renderer {
	tpl := new(renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func isHasChildren(data bson.M) int {
	result := 0
	if _, ok := data["child"]; ok {
		result = 1
	}

	return result
}

func (t *renderer) ReloadTemplates() {
	//t.template = template.Must(template.ParseGlob(t.location))
	newTpl := template.New("")
	err := filepath.Walk(t.location, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = newTpl.Funcs(template.FuncMap{"isHasChildren": isHasChildren}).ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	t.template = newTpl
}

func (t renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}
