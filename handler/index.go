package handler

import (
	"github.com/qiangxue/fasthttp-routing"
)

func IndexHandler(ctx *routing.Context) error {

	return RenderTemplate(ctx, "index.html", nil)
}
