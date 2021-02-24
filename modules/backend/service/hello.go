package service

import (
	"github.com/gin-gonic/gin"
	"hello-gin-micro/app"
	request "hello-gin-micro/modules/backend/form/request"
	Response "hello-gin-micro/modules/backend/form/response"
	"net/http"
)

type HelloService struct {
}

func (h *HelloService) SayHello(c *gin.Context, req request.SayHelloReq) Response.SayHelloRep {
	var res Response.SayHelloRep
	//请求参数验证 validate
	errValidateBool, errValidateMsg := app.GetError(req)
	if errValidateBool == false {
		res.Code = 401
		res.Msg = errValidateMsg
		return res
	}
	str := "你好" + req.Name + " 内容:" + req.Content
	res.Code = http.StatusOK
	res.Msg = "success"
	res.Data = str
	return res
}
