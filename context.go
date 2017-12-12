package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Store

type Store interface {
	Get(k interface{}) (v interface{}, exists bool)
	Set(k, v interface{})
}

func ToStore(c context.Context) Store {
	return c.(Store)
}

func GetRequest(ctx context.Context) *http.Request {
	return ctx.(*gin.Context).Request
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	return ctx.(*gin.Context).Writer
}

func Param(ctx context.Context, param string) string {
	return ctx.(*gin.Context).Param(param)
}

func AbortError(ctx context.Context, code int, err error) {
	ctx.(*gin.Context).AbortWithError(code, err)
}

func Redirect(ctx context.Context, code int, path string) {
	ctx.(*gin.Context).Redirect(code, path)
}
