package service

import (
	"encoding/json"
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"strings"
	"zhuaimao/models"
)

type GroupForm struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type User struct {
}

func (u User) Get(ctx *routing.Context) error {

	tye := ctx.Param("type")
	if tye == `group_add` {
		return RenderTemplate(ctx, `user_add.html`, `group`)
	} else if tye == `user_add` {
		return RenderTemplate(ctx, `user_add.html`, `user`)
	}

	return RenderTemplate(ctx, `user.html`, nil)
}

func (u User) One(ctx *routing.Context) error {
	return RenderTemplate(ctx, `user.html`, nil)
}

func (u User) Add(ctx *routing.Context) error {
	var err error

	tye := ctx.Param("type")

	if tye == "group" {
		err = u.AddGroup(ctx)
	} else if tye == "user" {
		err = u.AddUser(ctx)
	}

	if err != nil {
		return ErrorResponse(ctx, -1, err.Error())
	}

	return SuccessResponse(ctx, nil)
}

func (u User) AddUser(ctx *routing.Context) error {
	fmt.Println(string(ctx.Request.Body()))

	return nil
}

func (u User) AddGroup(ctx *routing.Context) error {
	groupModel := models.Group{}

	data := GroupForm{}
	err := json.Unmarshal(ctx.Request.Body(), &data)
	if err != nil {
		return err
	}

	return groupModel.Add(data.Name, strings.Join(data.Roles, ","))
}

func (u User) Update(ctx *routing.Context) error {
	return nil
}

func (u User) Delete(ctx *routing.Context) error {
	return nil
}
