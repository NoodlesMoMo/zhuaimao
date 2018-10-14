package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

var (
	permSrv = service.PermissionService{}
)

func PermissionHandler(ctx *routing.Context) error {

	switch string(ctx.Request.Header.Method()) {
	case `GET`:
		permSrv.Get(ctx)
	case `PUT`:
		permSrv.Add(ctx)
	}

	return nil
}
