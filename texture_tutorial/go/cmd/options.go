package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/plugins/go-texture-example/example_gif"
	"github.com/go-flutter-desktop/plugins/go-texture-example/example_image"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(1260, 580),
	flutter.AddPlugin(&example_image.ImagePlugin{}),
	flutter.AddPlugin(&example_gif.GifPlugin{}),
}
