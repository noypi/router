package router

import (
	"net/http"

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

	// http.HandlerFunc
	UseF(handlers ...http.HandlerFunc) IRoutesStd

	HandleF(httpMethod, relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	AnyF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	GETF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	POSTF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	DELETEF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	PATCHF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	PUTF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	OPTIONSF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd
	HEADF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd

	// http.Handler
	UseH(handlers ...http.Handler) IRoutesStd

	HandleH(httpMethod, relativePath string, handlers ...http.Handler) IRoutesStd
	AnyH(relativePath string, handlers ...http.Handler) IRoutesStd
	GETH(relativePath string, handlers ...http.Handler) IRoutesStd
	POSTH(relativePath string, handlers ...http.Handler) IRoutesStd
	DELETEH(relativePath string, handlers ...http.Handler) IRoutesStd
	PATCHH(relativePath string, handlers ...http.Handler) IRoutesStd
	PUTH(relativePath string, handlers ...http.Handler) IRoutesStd
	OPTIONSH(relativePath string, handlers ...http.Handler) IRoutesStd
	HEADH(relativePath string, handlers ...http.Handler) IRoutesStd

	// Handler
	UseC(handlers ...Handler) IRoutesStd

	HandleC(httpMethod, relativePath string, handlers ...Handler) IRoutesStd
	AnyC(relativePath string, handlers ...Handler) IRoutesStd
	GETC(relativePath string, handlers ...Handler) IRoutesStd
	POSTC(relativePath string, handlers ...Handler) IRoutesStd
	DELETEC(relativePath string, handlers ...Handler) IRoutesStd
	PATCHC(relativePath string, handlers ...Handler) IRoutesStd
	PUTC(relativePath string, handlers ...Handler) IRoutesStd
	OPTIONSC(relativePath string, handlers ...Handler) IRoutesStd
	HEADC(relativePath string, handlers ...Handler) IRoutesStd
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
	wrapMethod(e.Engine.Any, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) GET(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.GET, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) POST(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.POST, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) DELETE(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.DELETE, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) PATCH(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.PATCH, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) PUT(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.PUT, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) OPTIONS(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.OPTIONS, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) HEAD(relativePath string, handlers ...interface{}) IRoutesStd {
	wrapMethod(e.Engine.HEAD, relativePath, Wrap(handlers...))
	return e
}

func (e *EngineStd) Use(handlers ...interface{}) IRoutesStd {
	wrap(e.Engine.Use, Wrap(handlers...))
	return e
}

// http.Handler

func (e *EngineStd) HandleH(httpMethod, relativePath string, handlers ...http.Handler) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, WrapH(handlers...)...)
	return e
}

func (e *EngineStd) AnyH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.Any, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) GETH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.GET, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) POSTH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.POST, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) DELETEH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.DELETE, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) PATCHH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.PATCH, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) PUTH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.PUT, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) OPTIONSH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.OPTIONS, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) HEADH(relativePath string, handlers ...http.Handler) IRoutesStd {
	wrapMethod(e.Engine.HEAD, relativePath, WrapH(handlers...))
	return e
}

func (e *EngineStd) UseH(handlers ...http.Handler) IRoutesStd {
	wrap(e.Engine.Use, WrapH(handlers...))
	return e
}

// http.HandlerFunc

func (e *EngineStd) HandleF(httpMethod, relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, WrapF(handlers...)...)
	return e
}

func (e *EngineStd) AnyF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.Any, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) GETF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.GET, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) POSTF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.POST, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) DELETEF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.DELETE, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) PATCHF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.PATCH, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) PUTF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.PUT, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) OPTIONSF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.OPTIONS, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) HEADF(relativePath string, handlers ...http.HandlerFunc) IRoutesStd {
	wrapMethod(e.Engine.HEAD, relativePath, WrapF(handlers...))
	return e
}

func (e *EngineStd) UseF(handlers ...http.HandlerFunc) IRoutesStd {
	wrap(e.Engine.Use, WrapF(handlers...))
	return e
}

// Handler

func (e *EngineStd) HandleC(httpMethod, relativePath string, handlers ...Handler) IRoutesStd {
	e.Engine.Handle(httpMethod, relativePath, WrapC(handlers...)...)
	return e
}

func (e *EngineStd) AnyC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.Any, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) GETC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.GET, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) POSTC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.POST, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) DELETEC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.DELETE, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) PATCHC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.PATCH, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) PUTC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.PUT, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) OPTIONSC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.OPTIONS, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) HEADC(relativePath string, handlers ...Handler) IRoutesStd {
	wrapMethod(e.Engine.HEAD, relativePath, WrapC(handlers...))
	return e
}

func (e *EngineStd) UseC(handlers ...Handler) IRoutesStd {
	wrap(e.Engine.Use, WrapC(handlers...))
	return e
}
