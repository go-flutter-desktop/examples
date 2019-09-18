# texture_tutorial

The image texture plugin source code is stored in the same example: [./go-texture-example](./go-texture-example).  
The video texture plugin source code is available as a go-flutter plugin [video_player](https://github.com/go-flutter-desktop/plugins/tree/master/video_player).  

Multiples path are hard coded, run the app in the root of the flutter project!

The video example depends on [gmf](https://github.com/3d0c/gmf) (a Go FFmpeg
Bindings), please follow the Installation described in the gmf project README.

## Demo
<p align="center">
  <img src="https://user-images.githubusercontent.com/7476655/62157099-b1793180-b30c-11e9-9ca1-677f44432ebc.gif" width="650" align="center" alt="Demo of the
  example">
</p>

## :warning: Warning! :warning:
This plugin showcase can crash due to FFmpeg bindings errors (gmf).  
**For that reason, the video plugin in this example is off by default**  
If you want to give the video plugin example a shot, edit the
[go/cmd/options.go](https://github.com/go-flutter-desktop/examples/blob/master/texture_tutorial/go/cmd/options.go) file and add the `example_video.VideoPlugin` plugin.

```go
package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/plugins/go-texture-example/example_image"
	"github.com/go-flutter-desktop/plugins/video_player"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(1200, 500),
	flutter.AddPlugin(&video_player.VideoPlayerPlugin{}),
	flutter.AddPlugin(&example_image.ImagePlugin{}),
}

```

If you get errors when only using the `example_image.ImagePlugin` plugin, please report an issue!  
If you get errors when using both plugin, check if your system is compatible with the [video-to-goImage](https://github.com/3d0c/gmf/blob/f4b5acb7db5cbbda9a6209be1d0de5f552823f62/examples/video-to-goImage.go) gmf example.

#### note
User [@nzlov](https://github.com/nzlov) managed to build this example in `Parallels Desktop` by:  
Using [`msys2`](https://www.msys2.org/)  
And installing `pkg-config` using:
```
pacman -S  mingw-w64-x86_64-pkg-config 
```
[source](https://github.com/go-flutter-desktop/go-flutter/issues/248#issuecomment-529741140)
