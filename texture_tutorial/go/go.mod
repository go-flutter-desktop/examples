module texture_tutorial/go

go 1.15

require (
	github.com/go-flutter-desktop/go-flutter v0.52.1
	github.com/go-flutter-desktop/plugins/go-texture-example/example_gif v0.0.0-00010101000000-000000000000
	github.com/go-flutter-desktop/plugins/go-texture-example/example_image v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	golang.org/x/sys v0.0.0-20220627191245-f75cf1eec38b // indirect
)

replace github.com/go-flutter-desktop/plugins/go-texture-example/example_image => ../go-texture-example/image

replace github.com/go-flutter-desktop/plugins/go-texture-example/example_gif => ../go-texture-example/gif
