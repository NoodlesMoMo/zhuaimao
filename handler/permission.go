package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

func PermissionHandler(ctx *routing.Context) error {

	service.RenderTemplate(ctx, "permission.html", nil)

	return nil
}
