package handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var (
	router = routing.New()
)

func HandleRequest(ctx *fasthttp.RequestCtx) {
	router.HandleRequest(ctx)
}

func init() {
	authRouter := router.Group("")
	authRouter.Use(AuthMiddleHandler)
	authRouter.Get("/", IndexHandler)

	noAuthRouter := router.Group("")
	noAuthRouter.Get("/login", LoginHandler)
	noAuthRouter.Post("/sign-in", SignInHandler)

	noAuthRouter.Get("/static/*", StaticHandler)
}
