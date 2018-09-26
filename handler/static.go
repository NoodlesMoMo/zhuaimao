package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
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

func rootDir(sep ...string) string {
	pwd, _ := os.Getwd()

	pwds := []string{pwd}
	for _, s := range sep {
		pwds = append(pwds, s)
	}

	return path.Join(pwds...)
}
