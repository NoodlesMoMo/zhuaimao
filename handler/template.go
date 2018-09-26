package handler

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"html/template"
	"path"
	"path/filepath"
)

var (
	templates map[string]*template.Template
)

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templateBaseDir := rootDir("templates")

	bases, err := filepath.Glob(path.Join(templateBaseDir, "base/*.html"))
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(path.Join(templateBaseDir, "include/*.html"))

	for _, base := range bases {
		files := append(includes, base)
		templates[filepath.Base(base)] = template.Must(template.ParseFiles(files...))
	}
}

func RenderTemplate(ctx *routing.Context, name string, data interface{}) error {
	tpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s template not found", name)
	}
	ctx.SetContentType("text/html; charset=utf-8")

	return tpl.ExecuteTemplate(ctx.Response.BodyWriter(), name, data)
}
