module xbox_controller/go

go 1.15

require (
	github.com/go-flutter-desktop/go-flutter v0.52.1
	github.com/go-flutter-desktop/plugins/path_provider v0.4.0
	github.com/go-flutter-desktop/plugins/xbox v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	golang.org/x/sys v0.0.0-20220627191245-f75cf1eec38b // indirect
)

replace github.com/go-flutter-desktop/plugins/xbox => ../go-plugin-example/controller
