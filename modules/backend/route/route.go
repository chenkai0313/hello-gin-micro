package route

import (
	"github.com/gin-gonic/gin"
	"hello-gin-micro/modules/backend/controller"
	middleWare "hello-gin-micro/modules/backend/middleware"
)

func Route(r *gin.Engine) {
	authorized := r.Group("/backend")
	authorized.Use(middleWare.TracerWrapper) //链路追踪
	authorized.POST("/say-hello", controller.SayHello)

	//只需要token验证，不需要权限认证
	authorized.Use(middleWare.AuthMiddleware())
	{
	}
	//需要权限认证
	authorized.Use(middleWare.AuthPermissionMiddleware())
	{
	}
}
