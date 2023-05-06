package main

import (
	"github.com/livefir/fir"
)

func docsRoute() fir.RouteOptions {
	return fir.RouteOptions{
		fir.ID("livefir-docs"),
		fir.Content("docs.html"),
		fir.Partials("partials"),
		fir.Layout("docs_layout.html"),
	}
}
