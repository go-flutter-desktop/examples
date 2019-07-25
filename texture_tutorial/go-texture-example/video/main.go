package example_video

import (
	"flag"
	"fmt"
	"sync"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// VideoPlugin is a texture plugin example
type VideoPlugin struct {
	videoBuffer *ffmpegVideo
}

var _ flutter.PluginTexture = &VideoPlugin{} // compile-time type check

// InitPlugin is used because PluginTexture must implement flutter.Plugin
func (p *VideoPlugin) InitPlugin(messenger plugin.BinaryMessenger) (err error) {

	var srcFileName string
	// hard-coded path, run the app with `hover run` in the project root
	flag.StringVar(&srcFileName, "src", "SampleVideo_1280x720_10mb.mp4", "source video")
	flag.Parse()

	p.videoBuffer = &ffmpegVideo{}

	bufferSize := 30 // in frames
	return p.videoBuffer.Init(srcFileName, bufferSize)
}

// InitPluginTexture is used to create and manage backend textures
func (p *VideoPlugin) InitPluginTexture(registry *flutter.TextureRegistry) error {

	texture := registry.NewTexture()
	texture.ID = 2 // not a good practice, use platform message to align dart and golang textureId

	consumer := func(wg *sync.WaitGroup) { // start marking frame available
		imagePerSec := p.videoBuffer.GetFrameRate()
		start := time.Now()

		// TODO(drakirus): when the p.videoBuffer.Frames is empty, it doesn't
		// necessary means that the video is over, a frame might take time to be
		// encoded.
		// Step to reproduce: set a bufferSize to 1.
		for len(p.videoBuffer.Frames) != 0 {
			time.Sleep(time.Duration(imagePerSec*1000) * time.Millisecond)
			texture.FrameAvailable() // trigger p.textureHanler (display new frame)
		}

		fmt.Printf("texture_tutorial: video play time: %v, real time %vs\n", time.Since(start), p.videoBuffer.Duration())
		wg.Done()
	}

	go func() {
		fmt.Println("texture_tutorial: Start video")
		p.videoBuffer.Play(consumer) // on frame pending frames start consumer
		texture.UnRegister()
		texture.FrameAvailable() // repaint to clear scene
		p.videoBuffer.Free()
		fmt.Println("texture_tutorial: End video")
	}()

	return texture.Register(p.textureHanler)
}

// textureHanler is triggerd when texture.FrameAvailable is called
// TODO: scale the video down if the with and height of the flutter Widget is
// lower than the with and height of the video.
// This down-scaling will reduce memory usage.
// How to down-scale with golang gmf: https://github.com/3d0c/gmf/blob/22de4b5a3c28895fe1b0b787ab36d0b49e53375f/examples/scale_video.go#L36-L43
func (p *VideoPlugin) textureHanler(width, height int) (bool, *flutter.PixelBuffer) {

	vWidth, vHeight := p.videoBuffer.Bounds()

	// TODO(drakirus): Check if the video is over.
	// the rendering loop can be blocked if the video is over (no frame in
	// p.videoBuffer.Frames) and a texture.FrameAvailable is trigged by the
	// engine (when moving/resizing the window)
	pixels := <-p.videoBuffer.Frames // get the frame
	defer pixels.Free()

	return true, &flutter.PixelBuffer{ // send it to the scene
		Pix:    pixels.Data(),
		Width:  vWidth,
		Height: vHeight,
	}
}
