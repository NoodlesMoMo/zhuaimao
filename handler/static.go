package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"zhuaimao/service"
)

var (
	fsHandler = fasthttp.FSHandler(service.RootDir(), 0)
)

func StaticHandler(ctx *routing.Context) error {
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	fsHandler(ctx.RequestCtx)

	return nil
}
