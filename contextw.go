package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//type Context = gin.Context

type _ContextW struct {
	http.ResponseWriter
	c *gin.Context
}

func ContextW(w http.ResponseWriter) *gin.Context {
	return w.(*_ContextW).c
}

func NewContextW(c *gin.Context) *_ContextW {
	return &_ContextW{c: c}
}

// ResponseWriter

func (c *_ContextW) Header() http.Header {
	return c.c.Writer.Header()
}

func (c *_ContextW) Write(bb []byte) (int, error) {
	return c.c.Writer.Write(bb)
}

func (c *_ContextW) WriteHeader(n int) {
	c.c.Writer.WriteHeader(n)
}
