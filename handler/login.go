package handler

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
)

func LoginHandler(ctx *routing.Context) error {

	username, password := ctx.FormValue("username"), ctx.FormValue("password")

	fmt.Println("user_name:", username, "password:", password)

	return nil
}
