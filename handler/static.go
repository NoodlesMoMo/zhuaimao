package handler

import (
	"os"
	"path"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var (
	fsHandler = fasthttp.FSHandler(rootDir(), 0)
)

func StaticHandler(ctx *routing.Context) error {
	fsHandler(ctx.RequestCtx)
	return nil
}

func rootDir() string {
	pwd, _ := os.Getwd()

	return path.Join(pwd, "")
}
