package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

func LoginHandler(ctx *routing.Context) error {

	username, password := ctx.FormValue("username"), ctx.FormValue("password")

	if user, ok := service.CheckUser(username, password); ok {
		service.SetCookie(ctx, user)
		ctx.WriteString(`{"code": 0, "msg": "success", "url": "/"}`)
		return nil
	}

	ctx.WriteString(`{"code": -1, "msg": "failed"}`)

	return nil
}
