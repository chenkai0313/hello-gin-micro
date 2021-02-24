package context

import (
	"context"
	"github.com/gin-gonic/gin"
)

const contextTracerKey = "context-tracer"

func SetTracerContext(ginCtx *gin.Context, ctx context.Context) {
	ginCtx.Set(contextTracerKey, ctx)
}

func GetTracerContext(ginCtx *gin.Context) (ctx context.Context, exist bool) {
	v, exist := ginCtx.Get(contextTracerKey)
	if exist == false {
		ctx = context.TODO()
		return
	}

	ctx, exist = v.(context.Context)
	return
}
