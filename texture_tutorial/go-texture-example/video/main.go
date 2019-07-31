package example_video

import (
	"fmt"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// VideoPlugin is a texture plugin example
// Every frame of the video is extracted and send to flutter through the
// texture API. The process of extracting every frame is intensive, 4k video
// and high frame-rate video is sluggish.
type VideoPlugin struct {
	videoBuffer *ffmpegVideo

	// textureHanler can be called from 2 sources: texture.FrameAvailable and
	// flutter redraw (resize/move window).
	// when called by redraw, multiples new frames gets draw without proper timing,
	// leading to wrong frame rate. The video is accelerated.
	// This channel is used to keep ensure the timing is correct between goroutine.
	newFrame chan bool
}

var _ flutter.PluginTexture = &VideoPlugin{} // compile-time type check

// InitPlugin is used because PluginTexture must implement flutter.Plugin
func (p *VideoPlugin) InitPlugin(messenger plugin.BinaryMessenger) (err error) {

	p.newFrame = make(chan bool)
	p.videoBuffer = &ffmpegVideo{}

	bufferSize := 100 // in frames

	// hard-coded path, run the app with `hover run` in the project root
	return p.videoBuffer.Init("./SampleVideo_1280x720_10mb.mp4", bufferSize)
	// URL video also works!
	// return p.videoBuffer.Init("http://www.sample-videos.com/video123/mp4/720/big_buck_bunny_720p_20mb.mp4", bufferSize)
}

// InitPluginTexture is used to create and manage backend textures
func (p *VideoPlugin) InitPluginTexture(registry *flutter.TextureRegistry) error {

	texture := registry.NewTexture()
	texture.ID = 2 // not a good practice, use platform message to align dart and golang textureId

	consumer := func() {
		imagePerSec := p.videoBuffer.GetFrameRate()

		for p.videoBuffer.HasFrameAvailable() {
			time.Sleep(time.Duration(imagePerSec*1000) * time.Millisecond)
			p.newFrame <- true
			texture.FrameAvailable() // trigger p.textureHanler (display new frame)
		}
	}

	go func() {
		fmt.Println("texture_tutorial: Start video")
		start := time.Now()
		p.videoBuffer.Stream(consumer) // on pending frames, start consuming image in the channel
		fmt.Printf("texture_tutorial: video play time: %v, real time %vs\n", time.Since(start), p.videoBuffer.Duration())
		close(p.newFrame)
		texture.UnRegister()
		p.videoBuffer.Free()
		fmt.Println("texture_tutorial: End video")
	}()

	return texture.Register(p.textureHanler)
}

// textureHanler is triggerd when texture.FrameAvailable is called
//
// TODO: scale the video down if the with and height of the flutter Widget is
// lower than the with and height of the video.
// This down-scaling will reduce memory usage. (change the `SetWidth` and
// `SetHeight` of gmf.CodecCtx)
func (p *VideoPlugin) textureHanler(width, height int) (bool, *flutter.PixelBuffer) {

	vWidth, vHeight := p.videoBuffer.Bounds()

	if p.videoBuffer.Closed() {
		return false, nil
	}

	// Sync frames
	<-p.newFrame

	pixels := <-p.videoBuffer.Frames // get the frame, ! Block the main thread !
	defer pixels.Free()

	return true, &flutter.PixelBuffer{ // send the image to the scene
		Pix:    pixels.Data(),
		Width:  vWidth,
		Height: vHeight,
	}
}
