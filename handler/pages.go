package handler

import "github.com/qiangxue/fasthttp-routing"

func IndexPageHandler(ctx *routing.Context) error {
	data := "hello_templates"
	return RenderTemplate(ctx, "index.html", data)
}

func LoginPageHandler(ctx *routing.Context) error {
	return RenderTemplate(ctx, "login.html", nil)
}
