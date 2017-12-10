package router

import (
	"net/http"
)

func WrapF(handlers ...http.HandlerFunc) []HandlerFunc {
	ginh := make([]HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = ginWrapF(h)
	}
	return ginh
}

func WrapH(handlers ...http.Handler) []HandlerFunc {
	ginh := make([]HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = ginWrapH(h)
	}
	return ginh
}

func ginWrapF(f http.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		f(NewContextW(c), c.Request)
	}
}

func ginWrapH(h http.Handler) HandlerFunc {
	return func(c *Context) {
		h.ServeHTTP(NewContextW(c), c.Request)
	}
}

func wrapH(fn func(handlers ...HandlerFunc) IRoutes, handlers []http.Handler) {
	fn(WrapH(handlers...)...)
}

func wrapF(fn func(handlers ...HandlerFunc) IRoutes, handlers []http.HandlerFunc) {
	fn(WrapF(handlers...)...)
}

func wrapMethodH(fn func(relativePath string, handlers ...HandlerFunc) IRoutes, relativePath string, handlers []http.Handler) {
	fn(relativePath, WrapH(handlers...)...)
}

func wrapMethodF(fn func(relativePath string, handlers ...HandlerFunc) IRoutes, relativePath string, handlers []http.HandlerFunc) {
	fn(relativePath, WrapF(handlers...)...)
}
