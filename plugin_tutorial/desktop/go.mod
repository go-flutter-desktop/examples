module github.com/go-flutter-desktop/examples/plugin_tutorial/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.26.0
	github.com/go-flutter-desktop/plugins/go-plugin-example/battery v0.0.0-00010101000000-000000000000
	github.com/go-flutter-desktop/plugins/go-plugin-example/complex_data_structure v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
)

replace github.com/go-flutter-desktop/plugins/go-plugin-example/battery => ../go-plugin-example/battery

replace github.com/go-flutter-desktop/plugins/go-plugin-example/complex_data_structure => ../go-plugin-example/complex_data_structure
