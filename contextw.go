package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//type Context = gin.Context

type _ContextW struct {
	http.ResponseWriter
	*gin.Context

	keys map[interface{}]interface{}
}

func ContextW(w http.ResponseWriter) *gin.Context {
	return w.(*_ContextW).Context
}

func NewContextW(c *gin.Context) *_ContextW {
	return &_ContextW{Context: c}
}

func (c *_ContextW) Set(key, value interface{}) {
	if c.keys == nil {
		c.keys = make(map[interface{}]interface{})
	}
	c.keys[key] = value
}

func (c *_ContextW) Get(key interface{}) (value interface{}, exists bool) {
	if c.keys != nil {
		value, exists = c.keys[key]
	}
	return
}

// ResponseWriter

func (c *_ContextW) Header() http.Header {
	return c.Writer.Header()
}

func (c *_ContextW) Write(bb []byte) (int, error) {
	return c.Writer.Write(bb)
}

func (c *_ContextW) WriteHeader(n int) {
	c.Writer.WriteHeader(n)
}
