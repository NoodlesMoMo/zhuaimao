package handler

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

func UserHandler(ctx *routing.Context) error {

	fmt.Println(string(ctx.RequestCtx.Method()))

	RenderTemplate(ctx, `user.html`, nil)

	return nil
}
