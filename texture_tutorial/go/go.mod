module texture_tutorial/go

go 1.19

require (
	github.com/go-flutter-desktop/go-flutter v0.52.2
	github.com/go-flutter-desktop/plugins/go-texture-example/example_gif 01528a4714a5
	github.com/go-flutter-desktop/plugins/go-texture-example/example_image 01528a4714a5
	github.com/pkg/errors v0.9.1
)

require (
	github.com/Xuanwo/go-locale v1.1.0 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220712193148-63cf1f4ef61f // indirect
	golang.org/x/sys v0.0.0-20220627191245-f75cf1eec38b // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/go-flutter-desktop/plugins/go-texture-example/example_image => ../go-texture-example/image

replace github.com/go-flutter-desktop/plugins/go-texture-example/example_gif => ../go-texture-example/gif
