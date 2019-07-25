module github.com/go-flutter-desktop/examples/plugin_tutorial/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.26.1-0.20190724121833-53e362014bc5
	github.com/go-flutter-desktop/plugins/go-texture-example/example_image v0.0.1
	github.com/go-flutter-desktop/plugins/go-texture-example/example_video v0.0.1
	github.com/pkg/errors v0.8.1
	github.com/stretchr/objx v0.2.0 // indirect
)

replace github.com/go-flutter-desktop/plugins/go-texture-example/example_video => ../go-texture-example/video

replace github.com/go-flutter-desktop/plugins/go-texture-example/example_image => ../go-texture-example/image

// replace github.com/go-flutter-desktop/go-flutter => /home/drakirus/lab/go/src/github.com/go-flutter-desktop/go-flutter
