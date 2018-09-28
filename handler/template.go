package handler

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"html/template"
	"path"
	"path/filepath"
	"strings"
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

	baseParts := make(map[string][]string)
	for i := 0; i < len(bases); i++ {
		base := bases[i]
		seps := strings.SplitN(filepath.Base(base), "__", 2)
		if len(seps) > 1 {
			xb := seps[0]
			baseParts[xb] = append(baseParts[xb], base)
			bases = append(bases[:i], bases[i+1:]...)
			i--
		}
	}

	for _, base := range bases {
		files := append(includes, base)
		fb := filepath.Base(base)
		fbx := strings.TrimRight(fb, filepath.Ext(fb))
		if parts, ok := baseParts[fbx]; ok {
			files = append(files, parts...)
		}
		templates[fb] = template.Must(template.ParseFiles(files...))
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
