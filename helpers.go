package router

import (
	"github.com/gin-gonic/gin"
)

func DisableDebugging() {
	gin.SetMode(gin.ReleaseMode)
}

func EnableDebugging() {
	gin.SetMode(gin.DebugMode)
}
