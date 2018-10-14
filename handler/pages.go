package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

func IndexPageHandler(ctx *routing.Context) error {
	data := "hello_templates"
	return service.RenderTemplate(ctx, "index.html", data)
}

func LoginPageHandler(ctx *routing.Context) error {
	return service.RenderTemplate(ctx, "login.html", nil)
}

func MenuPageHandler(ctx *routing.Context) error {
	return service.RenderTemplate(ctx, "menu.html", nil)
}

func PermAddPageHandler(ctx *routing.Context) error {
	return service.RenderTemplate(ctx, "permission_add.html", nil)
}
