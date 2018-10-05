package handler

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

func UserHandler(ctx *routing.Context) error {

	id := ctx.Param("id")
	typ := ctx.Param("type")

	if typ == "" || typ == "list" {
		RenderTemplate(ctx, `user.html`, nil)
	}

	fmt.Print("method:", string(ctx.RequestCtx.Method()), " id:", id)

	return nil
}
