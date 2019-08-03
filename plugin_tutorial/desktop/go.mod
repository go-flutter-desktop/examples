module github.com/go-flutter-desktop/examples/plugin_tutorial/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.27.1
	github.com/go-flutter-desktop/plugins/go-plugin-example/battery v0.0.1
	github.com/go-flutter-desktop/plugins/go-plugin-example/complex v0.0.1
	github.com/pkg/errors v0.8.1
	github.com/stretchr/objx v0.2.0 // indirect
)

replace github.com/go-flutter-desktop/plugins/go-plugin-example/battery => ../go-plugin-example/battery

replace github.com/go-flutter-desktop/plugins/go-plugin-example/complex => ../go-plugin-example/complex
