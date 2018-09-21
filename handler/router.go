package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var (
	router = routing.New()
)

func HandleRequest(ctx *fasthttp.RequestCtx){
	router.HandleRequest(ctx)
}

func init() {
	authRouter := router.Group("/")
	authRouter.Use(AuthMiddleHandler)

	noAuthRouter := router.Group("/")
	noAuthRouter.To("GET,POS", "/login", LoginHandler)
}

