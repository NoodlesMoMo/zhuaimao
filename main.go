package main

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"sync"
)

func main() {

	waitGrp := sync.WaitGroup{}

	listener, err := GracefullListen(":8080")
	if err != nil {
		panic(err)
	}

	go func() {
		waitGrp.Add(1)
		defer waitGrp.Done()
		signalAction(listener)
	}()

	router := routing.New()

	router.Get("/data", func(ctx *routing.Context) error {
		ctx.WriteString("ooooooooookkkkkkkkkk")
		return nil
	})

	fasthttp.Serve(listener, router.HandleRequest)

	waitGrp.Wait()
}
