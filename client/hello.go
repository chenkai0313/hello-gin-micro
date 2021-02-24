package client

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"hello-gin-micro/app"
	"hello-gin-micro/lib/context"
	helloProto "hello-gin-micro/proto/hello"
	"time"
)

const Server_Name_Hello = "go.micro.srv.hello"

type HelloClient struct{}

//控制客户端超时时间方法
var Opss client.CallOption = func(o *client.CallOptions) {
	o.RequestTimeout = time.Second * 30
	o.DialTimeout = time.Second * 30
}

func (h HelloClient) HelloAdd(c *gin.Context, req *helloProto.AddData) (res *helloProto.Response, err error) {
	client := helloProto.NewHelloService(Server_Name_Hello, app.MicroService.Client())
	ctx, _ := context.GetTracerContext(c)
	r, err := client.Add(ctx, req)
	return r, err
}
