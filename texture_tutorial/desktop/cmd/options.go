package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/plugins/go-texture-example/example"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(400, 400),
	flutter.AddPlugin(&example.MyTexturePlugin{}),
}
