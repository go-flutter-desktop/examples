package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	controller "github.com/go-flutter-desktop/plugins/xbox"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(1600, 1100),
	flutter.WindowDimensionLimits(1600, 1100, 1600, 1100),
	flutter.AddPlugin(&controller.XBOXStream{}),
}
