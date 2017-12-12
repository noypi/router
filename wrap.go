package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler func(context.Context)

func WrapF(handlers ...http.HandlerFunc) []gin.HandlerFunc {
	ginh := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = GinWrapF(h)
	}
	return ginh
}

func WrapH(handlers ...http.Handler) []gin.HandlerFunc {
	ginh := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = GinWrapH(h)
	}
	return ginh
}

func WrapC(handlers ...Handler) []gin.HandlerFunc {
	ginh := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = GinWrapC(h)
	}
	return ginh
}

func Wrap(handlers ...interface{}) []gin.HandlerFunc {
	ginh := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = GinWrap(h)
	}
	return ginh
}

func GetRequest(ctx context.Context) *http.Request {
	return ctx.(*gin.Context).Request
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	return ctx.(*gin.Context).Writer
}

func GinWrap(f interface{}) gin.HandlerFunc {
	switch fn := f.(type) {
	case Handler:
		return GinWrapC(fn)
	case http.HandlerFunc:
		return GinWrapF(fn)
	case http.Handler:
		return GinWrapH(fn)
	default:
		panic("unsupported handler type")
	}
}

func GinWrapC(f Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c)
	}
}

func GinWrapF(f http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(NewContextW(c), c.Request)
	}
}

func GinWrapH(h http.Handler) gin.HandlerFunc {
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

func wrapC(fn func(handlers ...gin.HandlerFunc) gin.IRoutes, handlers []Handler) {
	fn(WrapC(handlers...)...)
}

func wrap(fn func(handlers ...gin.HandlerFunc) gin.IRoutes, handlers []interface{}) {
	fn(Wrap(handlers...)...)
}

func wrapMethod(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []interface{}) {
	fn(relativePath, Wrap(handlers...)...)
}

func wrapMethodC(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []Handler) {
	fn(relativePath, WrapC(handlers...)...)
}

func wrapMethodH(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []http.Handler) {
	fn(relativePath, WrapH(handlers...)...)
}

func wrapMethodF(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []http.HandlerFunc) {
	fn(relativePath, WrapF(handlers...)...)
}
