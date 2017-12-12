package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRequest(ctx context.Context) *http.Request {
	return ctx.(*gin.Context).Request
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	return ctx.(*gin.Context).Writer
}
