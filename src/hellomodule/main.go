package main

import (
	"expvar"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func sdsd(csd *expvar.Float) {

}

func main() {
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}
