package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"hello-gin-micro/app"
	"hello-gin-micro/lib"
	backendRoute "hello-gin-micro/modules/backend/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	//init config
	app.LoadConfig("./config")
	//init redis
	app.InitRedis()
	//init mysql
	app.InitMysql()

	//init log
	app.InitLogger()

	if os.Getenv("ENV") == "prod" {
		app.ZapLog.Info("server running", "pi已启动"+app.Config.Server.Port+" env prod")
	} else {
		app.ZapLog.Info("server running", "api已启动"+app.Config.Server.Port+" env test1")
	}

}

func main() {
	r := gin.New()

	//链路追踪
	t, closer, err := lib.NewJaegerTracer(app.Config.Server.Name, app.Config.Jaeger.Addr)
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(t)

	//建立微服务客户端
	app.MicroService = micro.NewService(
		micro.Name(app.Config.Server.Name),
		micro.Version("latest"),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(t)),
	)
	app.MicroService.Init()

	//后台路由模块
	backendRoute.Route(r)

	server := &http.Server{
		Addr:    ":" + app.Config.Server.Port,
		Handler: r,
	}
	//启动服务
	go server.ListenAndServe()

	gracefulExitWeb(server)
}

//抓取exit 信号 友好退出重启
func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL)
	sig := <-ch
	fmt.Println("got a signal", sig)
	cxt, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}
}
