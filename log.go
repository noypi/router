package router

import (
	"context"
	"log"

	"github.com/noypi/logfn"
)

type _logFuncType int

const (
	LogErrName _logFuncType = iota
	LogInfoName
	LogWarnName
	LogDebugName
)

func GetErrLog(ctx context.Context) logfn.LogFunc {
	return getLogFunc(ctx, LogErrName)
}

func GetInfoLog(ctx context.Context) logfn.LogFunc {
	return getLogFunc(ctx, LogInfoName)
}

func GetWarnLog(ctx context.Context) logfn.LogFunc {
	return getLogFunc(ctx, LogWarnName)
}

func GetDebugLog(ctx context.Context) logfn.LogFunc {
	return getLogFunc(ctx, LogDebugName)
}

func getLogFunc(ctx context.Context, name _logFuncType) (fn logfn.LogFunc) {
	if nil == ctx {
		return log.Printf
	}

	c := ToStore(ctx)

	if o, exists := c.Get(name); exists {
		fn = (o).(logfn.LogFunc)
	} else {
		fn = log.Printf
	}
	return
}
