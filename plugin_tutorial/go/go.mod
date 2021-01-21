module plugin_tutorial/go

go 1.15

require (
	github.com/go-flutter-desktop/go-flutter v0.42.0
	github.com/go-flutter-desktop/plugins/go-plugin-example/battery v0.0.0-00010101000000-000000000000
	github.com/go-flutter-desktop/plugins/go-plugin-example/complex v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
)

replace github.com/go-flutter-desktop/plugins/go-plugin-example/battery => ../go-plugin-example/battery

replace github.com/go-flutter-desktop/plugins/go-plugin-example/complex => ../go-plugin-example/complex
