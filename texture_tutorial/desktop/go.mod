module github.com/go-flutter-desktop/examples/plugin_tutorial/desktop

go 1.12

require (
	github.com/go-flutter-desktop/go-flutter v0.26.1-0.20190723223858-4cec5ef08ec4
	github.com/go-flutter-desktop/plugins/go-texture-example/example v0.0.0-00010101000000-000000000000
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7 // indirect
	github.com/pkg/errors v0.8.1
	github.com/stretchr/objx v0.2.0 // indirect
)

replace github.com/go-flutter-desktop/plugins/go-texture-example/example => ../go-texture-example/

// replace github.com/go-flutter-desktop/go-flutter => /home/drakirus/lab/go/src/github.com/go-flutter-desktop/go-flutter
