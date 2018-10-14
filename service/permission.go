package service

import (
	"encoding/json"
	"github.com/qiangxue/fasthttp-routing"
	"strconv"
	"strings"
	"zhuaimao/models"
)

type PermFormData struct {
	Name    string   `json:"name"`
	Slug    string   `json:"slug"`
	Methods []string `json:"methods"`
	Path    string   `json:"path"`
}

type PermissionService struct {
}

func (ps *PermissionService) Add(ctx *routing.Context) error {
	data := &PermFormData{}

	err := json.Unmarshal(ctx.Request.Body(), &data)
	if err != nil {
		return ErrorResponse(ctx, -1, err.Error())
	}

	model := models.Permission{}

	perm, err := model.Add(data.Name, data.Slug, strings.Join(data.Methods, ","), data.Path)
	if err != nil {
		return ErrorResponse(ctx, -2, err.Error())
	}

	return SuccessResponse(ctx, perm)
}

func (ps *PermissionService) Get(ctx *routing.Context) error {

	switch string(ctx.QueryArgs().Peek("type")) {
	case `list`:
		return ps.listPage(ctx)
	case `new`:
		return ps.newPage(ctx)
	}

	return nil
}

func (ps *PermissionService) listPage(ctx *routing.Context) error {
	model := models.Permission{}
	_page, _psize := string(ctx.QueryArgs().Peek("page")), string(ctx.QueryArgs().Peek("psize"))

	page, _ := strconv.Atoi(_page)
	psize, _ := strconv.Atoi(_psize)

	perms, err := model.List(page, psize)
	if err != nil {
		//TODO:
	}

	return RenderTemplate(ctx, "permission.html", perms)
}

func (ps *PermissionService) newPage(ctx *routing.Context) error {
	return RenderTemplate(ctx, "permission_add.html", nil)
}
