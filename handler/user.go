package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

func UserHandler(ctx *routing.Context) error {

	user := service.User{}

	switch string(ctx.RequestCtx.Method()) {
	case `GET`:
		id := ctx.Param("id")
		if id == "" {
			user.List(ctx)
		} else {
			user.One(ctx)
		}
	case `PUT`:
		user.Add(ctx)
	case `POST`:
	case `DELETE`:
	}
	user.One(ctx)

	return nil
}
