package controller

import (
	"github.com/gin-gonic/gin"
	request "hello-gin-micro/modules/backend/form/request"
	"hello-gin-micro/modules/backend/service"
	"net/http"
)

var helloService service.HelloService

func SayHello(c *gin.Context) {
	var req request.SayHelloReq
	_ = c.ShouldBind(&req)
	res := helloService.SayHello(c, req)
	c.JSON(http.StatusOK, res)
}
