package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRoutesStd interface {
	UseH(...http.Handler) IRoutesStd
	UseF(...http.HandlerFunc) IRoutesStd
	UseC(...Handler) IRoutesStd
	Use(...interface{}) IRoutesStd

	Handle(string, string, ...interface{}) IRoutesStd
	Any(string, ...interface{}) IRoutesStd
	GET(string, ...interface{}) IRoutesStd
	POST(string, ...interface{}) IRoutesStd
	DELETE(string, ...interface{}) IRoutesStd
	PATCH(string, ...interface{}) IRoutesStd
	PUT(string, ...interface{}) IRoutesStd
	OPTIONS(string, ...interface{}) IRoutesStd
	HEAD(string, ...interface{}) IRoutesStd

	HandleC(string, string, ...Handler) IRoutesStd
	AnyC(string, ...Handler) IRoutesStd
	GETC(string, ...Handler) IRoutesStd
	POSTC(string, ...Handler) IRoutesStd
	DELETEC(string, ...Handler) IRoutesStd
	PATCHC(string, ...Handler) IRoutesStd
	PUTC(string, ...Handler) IRoutesStd
	OPTIONSC(string, ...Handler) IRoutesStd
	HEADC(string, ...Handler) IRoutesStd

	HandleF(string, string, ...http.HandlerFunc) IRoutesStd
	AnyF(string, ...http.HandlerFunc) IRoutesStd
	GETF(string, ...http.HandlerFunc) IRoutesStd
	POSTF(string, ...http.HandlerFunc) IRoutesStd
	DELETEF(string, ...http.HandlerFunc) IRoutesStd
	PATCHF(string, ...http.HandlerFunc) IRoutesStd
	PUTF(string, ...http.HandlerFunc) IRoutesStd
	OPTIONSF(string, ...http.HandlerFunc) IRoutesStd
	HEADF(string, ...http.HandlerFunc) IRoutesStd

	HandleH(string, string, ...http.Handler) IRoutesStd
	AnyH(string, ...http.Handler) IRoutesStd
	GETH(string, ...http.Handler) IRoutesStd
	POSTH(string, ...http.Handler) IRoutesStd
	DELETEH(string, ...http.Handler) IRoutesStd
	PATCHH(string, ...http.Handler) IRoutesStd
	PUTH(string, ...http.Handler) IRoutesStd
	OPTIONSH(string, ...http.Handler) IRoutesStd
	HEADH(string, ...http.Handler) IRoutesStd
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

// func(context.Context)
func (e *EngineStd) HandleC(httpMethod, relativePath string, handlers ...Handler) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, WrapC(handlers...)...)
	return e
}

func (e *EngineStd) AnyC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.Any, relativePath, handlers)
	return e
}

func (e *EngineStd) GETC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.GET, relativePath, handlers)
	return e
}

func (e *EngineStd) POSTC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.POST, relativePath, handlers)
	return e
}

func (e *EngineStd) DELETEC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.DELETE, relativePath, handlers)
	return e
}

func (e *EngineStd) PATCHC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.PATCH, relativePath, handlers)
	return e
}

func (e *EngineStd) PUTC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.PUT, relativePath, handlers)
	return e
}

func (e *EngineStd) OPTIONSC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.OPTIONS, relativePath, handlers)
	return e
}

func (e *EngineStd) HEADC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethodC(e.Engine.HEAD, relativePath, handlers)
	return e
}

// with method
func (e *EngineStd) NoMethodH(handlers ...http.Handler) {
	e.Engine.NoMethod(WrapH(handlers...)...)
}

func (e *EngineStd) NoMethod(handlers ...http.HandlerFunc) {
	e.Engine.NoMethod(WrapF(handlers...)...)
}

// with method
func (e *EngineStd) HandleH(httpMethod, relativePath string, handlers ...http.Handler) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, WrapH(handlers...)...)
	return e
}

func (e *EngineStd) HandleF(httpMethod, relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, WrapF(handlers...)...)
	return e
}

func (e *EngineStd) AnyH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.Any, relativePath, handlers)
	return e
}

func (e *EngineStd) AnyF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.Any, relativePath, handlers)
	return e
}

func (e *EngineStd) GETH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.GET, relativePath, handlers)
	return e
}

func (e *EngineStd) GETF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.GET, relativePath, handlers)
	return e
}

func (e *EngineStd) POSTH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.POST, relativePath, handlers)
	return e
}

func (e *EngineStd) POSTF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.POST, relativePath, handlers)
	return e
}

func (e *EngineStd) DELETEH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.DELETE, relativePath, handlers)
	return e
}

func (e *EngineStd) DELETEF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.DELETE, relativePath, handlers)
	return e
}

func (e *EngineStd) PATCHH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.PATCH, relativePath, handlers)
	return e
}

func (e *EngineStd) PATCHF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.PATCH, relativePath, handlers)
	return e
}

func (e *EngineStd) PUTH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.PUT, relativePath, handlers)
	return e
}

func (e *EngineStd) PUTF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.PUT, relativePath, handlers)
	return e
}

func (e *EngineStd) OPTIONSH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.OPTIONS, relativePath, handlers)
	return e
}

func (e *EngineStd) OPTIONSF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.OPTIONS, relativePath, handlers)
	return e
}

func (e *EngineStd) HEADH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.HEAD, relativePath, handlers)
	return e
}

func (e *EngineStd) HEADF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.HEAD, relativePath, handlers)
	return e
}

// without method

func (e *EngineStd) UseH(handlers ...http.Handler) IRoutesStd {
	wrapH(e.Engine.Use, handlers)
	return e
}

func (e *EngineStd) UseF(handlers ...http.HandlerFunc) IRoutesStd {
	wrapF(e.Engine.Use, handlers)
	return e
}

func (e *EngineStd) UseC(handlers ...Handler) IRoutesStd {
	wrapC(e.Engine.Use, handlers)
	return e
}

func (e *EngineStd) Use(handlers ...interface{}) IRoutesStd {
	wrap(e.Engine.Use, handlers)
	return e
}
