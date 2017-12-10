package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WrapF(handlers ...http.HandlerFunc) []gin.HandlerFunc {
	ginh := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = ginWrapF(h)
	}
	return ginh
}

func WrapH(handlers ...http.Handler) []gin.HandlerFunc {
	ginh := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = ginWrapH(h)
	}
	return ginh
}

func GetRequest(ctx context.Context) *http.Request {
	return ctx.(*gin.Context).Request
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	return ctx.(*gin.Context).Writer
}

func ginWrapF(f http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(NewContextW(c), c.Request)
	}
}

func ginWrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(NewContextW(c), c.Request)
	}
}

func wrapH(fn func(handlers ...gin.HandlerFunc) gin.IRoutes, handlers []http.Handler) {
	fn(WrapH(handlers...)...)
}

func wrapF(fn func(handlers ...gin.HandlerFunc) gin.IRoutes, handlers []http.HandlerFunc) {
	fn(WrapF(handlers...)...)
}

func wrapMethodH(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []http.Handler) {
	fn(relativePath, WrapH(handlers...)...)
}

func wrapMethodF(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []http.HandlerFunc) {
	fn(relativePath, WrapF(handlers...)...)
}
