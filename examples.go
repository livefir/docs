package main

import (
	"github.com/livefir/fir"
)

func examplesRoute() fir.RouteOptions {
	return fir.RouteOptions{
		fir.ID("livefir-examples"),
		fir.Content("examples.html"),
		fir.Partials("partials"),
		fir.Layout("examples_layout.html"),
	}
}
