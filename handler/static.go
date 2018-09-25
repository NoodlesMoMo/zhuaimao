package handler

import (
	"github.com/valyala/fasthttp"
	"github.com/qiangxue/fasthttp-routing"
	"os"
	"path"
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
