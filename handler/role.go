package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

func RoleHandler(ctx *routing.Context) error {

	service.RenderTemplate(ctx, "role.html", nil)

	return nil
}
