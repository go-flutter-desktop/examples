module github.com/go-flutter-desktop/examples/plugin_tutorial/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.36.2-0.20200402135958-a526e1d973c5
	github.com/go-flutter-desktop/plugins/go-plugin-example/battery v0.0.1
	github.com/go-flutter-desktop/plugins/go-plugin-example/complex v0.0.1
	github.com/pkg/errors v0.9.1
)

replace github.com/go-flutter-desktop/plugins/go-plugin-example/battery => ../go-plugin-example/battery

replace github.com/go-flutter-desktop/plugins/go-plugin-example/complex => ../go-plugin-example/complex
