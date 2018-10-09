package service

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

type User struct {
}

func (u User) List(ctx *routing.Context) error {
	return RenderTemplate(ctx, `user.html`, nil)
}

func (u User) One(ctx *routing.Context) error {
	return RenderTemplate(ctx, `user.html`, nil)
}

func (u User) Add(ctx *routing.Context) error {
	fmt.Println(string(ctx.Request.Body()))

	return nil
}

func (u User) Update(ctx *routing.Context) error {
	return nil
}

func (u User) Delete(ctx *routing.Context) error {
	return nil
}
