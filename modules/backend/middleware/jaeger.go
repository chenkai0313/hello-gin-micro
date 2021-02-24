package middleware

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	libContext "hello-gin-micro/lib/context"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
)

const contextTracerKey = "Tracer-context"

// sf sampling frequency
var sf = 100

func init() {
	rand.Seed(time.Now().Unix())
}

// TracerWrapper tracer 中间件
func TracerWrapper(c *gin.Context) {
	sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path)
	tracer := opentracing.GlobalTracer()
	md := make(map[string]string)
	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if err == nil {
		sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
		tracer = sp.Tracer()
	}

	defer sp.Finish()

	if err := tracer.Inject(sp.Context(),
		opentracing.TextMap,
		opentracing.TextMapCarrier(md)); err != nil {
		log.Error(err)
		//log.Log(err)
	}

	ctx := context.TODO()
	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	libContext.SetTracerContext(c, ctx)

	c.Next()
}
