package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo"
)

// Template ...
type Template struct {
	templates map[string]*template.Template
}

// Edges ...
type Edges []struct {
	Node Node
}

// Node ...
type Node struct {
	Text string
}

// Add ...
func (t Template) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	t.templates[name] = tmpl
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if _, ok := t.templates[name]; ok == false {
		// not such view
		return fmt.Errorf("no such view. (%s)", name)
	}
	return t.templates[name].Execute(w, data)
}

// NewTemplate creates a new template
func NewTemplate(templatesDir string) *Template {
	ext := ".html"
	ins := Template{
		templates: map[string]*template.Template{},
	}

	layout := templatesDir + "/layouts/base" + ext

	_, err := os.Stat(layout)
	if err != nil {
		log.Panicf("cannot find %s", layout)
		os.Exit(1)
	}

	partials, err := filepath.Glob(templatesDir + "partials/" + "*" + ext)
	if err != nil {
		log.Print("cannot find " + templatesDir + "partials/" + "*" + ext)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"caption": func(edge Edges) template.HTML {
			return template.HTML("edge[0].Node.Text")
		},
		"safehtml": func(text string) template.HTML {
			return template.HTML(text)
		},
		"unixToDate": func(value int64) string {
			tm := time.Unix(value, 0)

			return fmt.Sprintf("%s", tm.Format("2006-01-02"))
		},
		"unixToDatetime": func(value int) string {
			tm := time.Unix(int64(value), 0)

			return fmt.Sprintf("%s", tm.Format("2006-01-02-03:04:05"))
		},
	}

	views, _ := filepath.Glob(templatesDir + "**/*" + ext)
	for _, view := range views {
		dir, file := filepath.Split(view)
		dir = strings.Replace(dir, templatesDir, "", 1)
		file = strings.TrimSuffix(file, ext)
		renderName := dir + file

		tmplfiles := append([]string{layout, view}, partials...)
		tmpl := template.Must(template.New(filepath.Base(layout)).Funcs(funcMap).ParseFiles(tmplfiles...))
		ins.Add(renderName, tmpl)
		log.Printf("renderName: %s, layout: %s", renderName, layout)
	}
	return &ins
}
