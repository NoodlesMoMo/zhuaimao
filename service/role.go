package service

import (
	"encoding/json"
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"strconv"
	"zhuaimao/models"
)

type RoleService struct {
}

func (r *RoleService) Add(ctx *routing.Context) error {
	roleModel := models.Role{}

	body := ctx.Request.Body()

	if len(body) == 0 {
		return ErrorResponse(ctx, -1, "invalid param")
	}

	data := make(map[string]string)
	err := json.Unmarshal(body, &data)
	if err != nil {
		return ErrorResponse(ctx, -2, err.Error())
	}

	role, err := roleModel.Add(data["role_name"])
	if err != nil {
		return ErrorResponse(ctx, -3, err.Error())
	}

	SuccessResponse(ctx, &role)

	return nil
}

func (r *RoleService) Cat(ctx *routing.Context) error {

	model := models.Role{}

	_page := string(ctx.QueryArgs().Peek("page"))
	_psize := string(ctx.QueryArgs().Peek("size"))

	page, _ := strconv.Atoi(_page)
	psize, _ := strconv.Atoi(_psize)

	result, err := model.List(page, psize)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return RenderTemplate(ctx, "role.html", result)
}
