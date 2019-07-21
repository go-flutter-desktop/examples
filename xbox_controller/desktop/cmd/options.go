package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	controller "github.com/go-flutter-desktop/plugins/xbox"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 500),
	flutter.WindowDimensionLimits(800, 500, 800, 500),
	flutter.AddPlugin(&controller.XBOXStream{}),
}
