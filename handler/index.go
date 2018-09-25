package handler

import "github.com/qiangxue/fasthttp-routing"

func IndexHandler(ctx *routing.Context) error {
	ctx.SetContentType("text/html")

	ctx.WriteString("<h1>Hello</h1>")

	return nil
}
