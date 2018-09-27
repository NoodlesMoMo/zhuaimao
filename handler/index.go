package handler

import (
	"github.com/qiangxue/fasthttp-routing"
)

func IndexHandler(ctx *routing.Context) error {

	data := "hello_templates"

	return RenderTemplate(ctx, "index.html", data)
}
