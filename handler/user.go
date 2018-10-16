package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

var (
	userSrv = service.User{}
)

func UserHandler(ctx *routing.Context) error {

	var err error

	switch string(ctx.RequestCtx.Method()) {
	case `GET`:
		err = userSrv.Get(ctx)
	case `PUT`:
		err = userSrv.Add(ctx)
	case `POST`:
	case `DELETE`:
	}

	return err
}
