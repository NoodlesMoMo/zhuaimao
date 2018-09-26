package handler

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var (
	indexFileHandler = fasthttp.FSHandler(rootDir("static/"), 1)
)

func IndexHandler(ctx *routing.Context) error {

	fmt.Println("path:", string(ctx.Request.RequestURI()))

	indexFileHandler(ctx.RequestCtx)

	return nil
}
