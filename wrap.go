package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler func(context.Context)

func Wrap(handlers ...interface{}) []gin.HandlerFunc {
	ginh := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = GinWrap(h)
	}
	return ginh
}

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

func GinWrap(f interface{}) gin.HandlerFunc {
	switch fn := f.(type) {

	case func(http.ResponseWriter, *http.Request):
		return GinWrapF(fn)
	case http.HandlerFunc:
		return GinWrapF(fn)

	case func(context.Context):
		return GinWrapC(fn)
	case Handler:
		return GinWrapC(fn)

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

func wrap(fn func(handlers ...gin.HandlerFunc) gin.IRoutes, handlers []gin.HandlerFunc) {
	fn(handlers...)
}

func wrapMethod(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []gin.HandlerFunc) {
	fn(relativePath, handlers...)
}
