module github.com/go-flutter-desktop/examples/xbox_controller/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.35.0
	github.com/go-flutter-desktop/plugins/xbox v0.0.1
	github.com/pkg/errors v0.9.1
)

replace github.com/go-flutter-desktop/plugins/xbox => ../go-plugin-example/controller
