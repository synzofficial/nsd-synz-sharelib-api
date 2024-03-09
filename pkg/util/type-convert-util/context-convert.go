package typeconvertutil

import (
	"context"

	"github.com/gin-gonic/gin"
)

func ConvertGinContext(ctx context.Context) context.Context {
	if c, isGinContext := ctx.(*gin.Context); isGinContext {
		return c.Request.Context()
	}
	return ctx
}
