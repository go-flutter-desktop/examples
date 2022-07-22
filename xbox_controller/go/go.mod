module xbox_controller/go

go 1.18

require (
	github.com/go-flutter-desktop/go-flutter v0.52.1
	github.com/go-flutter-desktop/plugins/path_provider v0.5.0
	github.com/go-flutter-desktop/plugins/xbox v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
)

require (
	github.com/Xuanwo/go-locale v1.1.0 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220712193148-63cf1f4ef61f // indirect
	golang.org/x/sys v0.0.0-20220627191245-f75cf1eec38b // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/go-flutter-desktop/plugins/xbox => ../go-plugin-example/controller
