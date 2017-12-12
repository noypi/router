package router

import (
	"github.com/gin-gonic/gin"
)

type IRoutesStd interface {

	// supported handlers are
	//   - func(w http.ResponseWriter, r *http.Request)
	//   - http.Handler
	//   - func(c context.Context)

	Use(handlers ...interface{}) IRoutesStd

	Handle(httpMethod, relativePath string, handlers ...interface{}) IRoutesStd
	Any(relativePath string, handlers ...interface{}) IRoutesStd
	GET(relativePath string, handlers ...interface{}) IRoutesStd
	POST(relativePath string, handlers ...interface{}) IRoutesStd
	DELETE(relativePath string, handlers ...interface{}) IRoutesStd
	PATCH(relativePath string, handlers ...interface{}) IRoutesStd
	PUT(relativePath string, handlers ...interface{}) IRoutesStd
	OPTIONS(relativePath string, handlers ...interface{}) IRoutesStd
	HEAD(relativePath string, handlers ...interface{}) IRoutesStd
}

//type IRoutes = gin.IRoutes
//type HandlerFunc = gin.HandlerFunc

type EngineStd struct {
	*gin.Engine
}

func New() *EngineStd {
	return &EngineStd{gin.New()}
}

func WrapGin(o *gin.Engine) *EngineStd {
	return &EngineStd{o}
}

func (e *EngineStd) Gin() *gin.Engine {
	return e.Engine
}

// func (interface{})
func (e *EngineStd) Handle(httpMethod, relativePath string, handlers ...interface{}) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, Wrap(handlers...)...)
	return e
}

func (e *EngineStd) Any(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.Any, relativePath, handlers)
	return e
}

func (e *EngineStd) GET(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.GET, relativePath, handlers)
	return e
}

func (e *EngineStd) POST(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.POST, relativePath, handlers)
	return e
}

func (e *EngineStd) DELETE(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.DELETE, relativePath, handlers)
	return e
}

func (e *EngineStd) PATCH(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.PATCH, relativePath, handlers)
	return e
}

func (e *EngineStd) PUT(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.PUT, relativePath, handlers)
	return e
}

func (e *EngineStd) OPTIONS(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.OPTIONS, relativePath, handlers)
	return e
}

func (e *EngineStd) HEAD(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.HEAD, relativePath, handlers)
	return e
}

func (e *EngineStd) Use(handlers ...interface{}) IRoutesStd {
	wrap(e.Engine.Use, handlers)
	return e
}
