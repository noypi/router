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

func GinWrap(f interface{}) gin.HandlerFunc {
	switch fn := f.(type) {
		
	case func(http.ResponseWriter, r *http.Request):
		fallthrough
	case http.HandlerFunc:
		return GinWrapF(fn)

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

func wrap(fn func(handlers ...gin.HandlerFunc) gin.IRoutes, handlers []interface{}) {
	fn(Wrap(handlers...)...)
}

func wrapMethod(fn func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers []interface{}) {
	fn(relativePath, Wrap(handlers...)...)
}
