package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context = gin.Context

type _ContextW struct {
	http.ResponseWriter
	*Context
}

func (c *_ContextW) Header() http.Header {
	return c.Context.Writer.Header()
}

func ContextW(w http.ResponseWriter) *Context {
	return w.(*_ContextW).Context
}

func NewContextW(c *Context) *_ContextW {
	return &_ContextW{
		Context:        c,
		ResponseWriter: c.Writer,
	}
}
