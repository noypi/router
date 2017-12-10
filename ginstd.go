package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRoutesStd interface {
	UseH(...http.Handler) IRoutesStd
	Use(...http.HandlerFunc) IRoutesStd

	HandleH(string, string, ...http.Handler) IRoutesStd
	AnyH(string, ...http.Handler) IRoutesStd
	GETH(string, ...http.Handler) IRoutesStd
	POSTH(string, ...http.Handler) IRoutesStd
	DELETEH(string, ...http.Handler) IRoutesStd
	PATCHH(string, ...http.Handler) IRoutesStd
	PUTH(string, ...http.Handler) IRoutesStd
	OPTIONSH(string, ...http.Handler) IRoutesStd
	HEADH(string, ...http.Handler) IRoutesStd

	Handle(string, string, ...http.HandlerFunc) IRoutesStd
	Any(string, ...http.HandlerFunc) IRoutesStd
	GET(string, ...http.HandlerFunc) IRoutesStd
	POST(string, ...http.HandlerFunc) IRoutesStd
	DELETE(string, ...http.HandlerFunc) IRoutesStd
	PATCH(string, ...http.HandlerFunc) IRoutesStd
	PUT(string, ...http.HandlerFunc) IRoutesStd
	OPTIONS(string, ...http.HandlerFunc) IRoutesStd
	HEAD(string, ...http.HandlerFunc) IRoutesStd
}

type IRoutes = gin.IRoutes
type HandlerFunc = gin.HandlerFunc

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

func WrapF(handlers ...http.HandlerFunc) []HandlerFunc {
	ginh := make([]HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = gin.WrapF(h)
	}
	return ginh
}

func WrapH(handlers ...http.Handler) []HandlerFunc {
	ginh := make([]HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginh[i] = gin.WrapH(h)
	}
	return ginh
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

func (e *EngineStd) Handle(httpMethod, relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, WrapF(handlers...)...)
	return e
}

func (e *EngineStd) AnyH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.Any, relativePath, handlers)
	return e
}

func (e *EngineStd) Any(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.Any, relativePath, handlers)
	return e
}

func (e *EngineStd) GETH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.GET, relativePath, handlers)
	return e
}

func (e *EngineStd) GET(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.GET, relativePath, handlers)
	return e
}

func (e *EngineStd) POSTH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.POST, relativePath, handlers)
	return e
}

func (e *EngineStd) POST(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.POST, relativePath, handlers)
	return e
}

func (e *EngineStd) DELETEH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.DELETE, relativePath, handlers)
	return e
}

func (e *EngineStd) DELETE(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.DELETE, relativePath, handlers)
	return e
}

func (e *EngineStd) PATCHH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.PATCH, relativePath, handlers)
	return e
}

func (e *EngineStd) PATCH(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.PATCH, relativePath, handlers)
	return e
}

func (e *EngineStd) PUTH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.PUT, relativePath, handlers)
	return e
}

func (e *EngineStd) PUT(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.PUT, relativePath, handlers)
	return e
}

func (e *EngineStd) OPTIONSH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.OPTIONS, relativePath, handlers)
	return e
}

func (e *EngineStd) OPTIONS(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.OPTIONS, relativePath, handlers)
	return e
}

func (e *EngineStd) HEADH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethodH(e.Engine.HEAD, relativePath, handlers)
	return e
}

func (e *EngineStd) HEAD(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethodF(e.Engine.HEAD, relativePath, handlers)
	return e
}

// without method

func (e *EngineStd) UseH(handlers ...http.Handler) IRoutesStd {
	wrapH(e.Engine.Use, handlers)
	return e
}

func (e *EngineStd) Use(handlers ...http.HandlerFunc) IRoutesStd {
	wrapF(e.Engine.Use, handlers)
	return e
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
