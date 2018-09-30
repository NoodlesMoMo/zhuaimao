package handler

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

func UserHandler(ctx *routing.Context) error {

	id := ctx.Param("id")

	fmt.Print("method:", string(ctx.RequestCtx.Method()), " id:", id)

	RenderTemplate(ctx, `user.html`, nil)

	return nil
}
