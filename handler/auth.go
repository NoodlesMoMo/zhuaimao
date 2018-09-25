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
