package handler

import (
	"fmt"
	"zhuaimao/models"
	"zhuaimao/service"

	"github.com/qiangxue/fasthttp-routing"
)

func AuthMiddleHandler(ctx *routing.Context) error {
	cookie := string(ctx.Request.Header.Cookie(service.IMECookieKey)[:])

	id := models.InitSession(ctx).GetUserId(cookie)
	fmt.Println(id)

	return ctx.Next()
}

func SignInHandler(ctx *routing.Context) error {

	username, password := ctx.FormValue("username"), ctx.FormValue("password")
	if user, ok := service.CheckUser(username, password); ok {
		service.SetCookie(ctx, user)
		ctx.WriteString(`{"code": 0, "msg": "success", "url": "/"}`)
		return nil
	}

	ctx.WriteString(`{"code": -1, "msg": "failed"}`)

	return nil
}
