module github.com/go-flutter-desktop/examples/plugin_tutorial/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.33.0
	github.com/go-flutter-desktop/plugins/go-texture-example/example_image v0.0.1
	github.com/go-flutter-desktop/plugins/video_player v0.0.2
	github.com/pkg/errors v0.8.1
)

replace github.com/go-flutter-desktop/plugins/go-texture-example/example_image => ../go-texture-example/image

// replace github.com/go-flutter-desktop/go-flutter => /home/drakirus/lab/go/src/github.com/go-flutter-desktop/go-flutter
