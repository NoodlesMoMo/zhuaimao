package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"zhuaimao/service"
)

var (
	roleSrv = service.RoleService{}
)

func RoleHandler(ctx *routing.Context) error {

	switch string(ctx.Request.Header.Method()) {
	case `PUT`:
		roleSrv.Add(ctx)
	case `GET`:
		roleSrv.Cat(ctx)
	}

	return nil
}
