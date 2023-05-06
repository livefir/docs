package main

import (
	"net/http"
	"sync/atomic"

	"github.com/livefir/fir"
)

func indexRoute() fir.RouteOptions {
	var count int32
	return fir.RouteOptions{
		fir.ID("livefir-index"),
		fir.Content("app.html"),
		fir.Layout("layout.html"),
		fir.Partials("partials"),
		fir.OnLoad(func(ctx fir.RouteContext) error {
			return ctx.KV("count", atomic.LoadInt32(&count))
		}),
		fir.OnEvent("inc", func(ctx fir.RouteContext) error {
			return ctx.KV("count", atomic.AddInt32(&count, 1))
		}),
		fir.OnEvent("dec", func(ctx fir.RouteContext) error {
			return ctx.KV("count", atomic.AddInt32(&count, -1))
		}),
	}
}

func main() {
	controller := fir.NewController("livefir-docs", fir.DevelopmentMode(true))
	http.Handle("/", controller.RouteFunc(indexRoute))
	http.Handle("/docs", controller.RouteFunc(docsRoute))
	http.Handle("/examples", controller.RouteFunc(examplesRoute))
	http.ListenAndServe(":8080", nil)
}
