package main

import (
	"github.com/valyala/fasthttp"
	"sync"
	"zhuaimao/handler"
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

	fasthttp.Serve(listener, handler.HandleRequest)

	waitGrp.Wait()
}
