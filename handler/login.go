package handler

import (
	"github.com/qiangxue/fasthttp-routing"
)

func LoginHandler(ctx *routing.Context) error {
	return RenderTemplate(ctx, "login.html", nil)
}
