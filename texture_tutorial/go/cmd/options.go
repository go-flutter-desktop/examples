package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/plugins/go-texture-example/example_image"
	// "github.com/go-flutter-desktop/plugins/video_player"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(1200, 500),
	// flutter.AddPlugin(&video_player.VideoPlayerPlugin{}),
	flutter.AddPlugin(&example_image.ImagePlugin{}),
}
