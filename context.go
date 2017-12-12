package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
