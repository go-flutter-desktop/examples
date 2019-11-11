package example_gif

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// GifPlugin is a texture plugin example
type GifPlugin struct {
	gif  *gif.GIF
	rgba *image.RGBA // in-memory buffer for the image
}

var _ flutter.PluginTexture = &GifPlugin{} // compile-time type check

// InitPlugin is used because PluginTexture must implement flutter.Plugin.
// InitPlugin block the first frame, it's better to have the image decoded
// on demand in a HandleFunc (non-bloking).
func (p *GifPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {

	// hard-coded path, run the app with `hover run` in the project root
	file := "./SampleVideo.gif"
	gifFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("gif %q not found on disk: %v\n", file, err)
		return err
	}

	p.gif, err = gif.DecodeAll(gifFile)
	if err != nil {
		fmt.Printf("error decoding file %s:%v\n", file, err)
		return err
	}

	gifWidth, gifHeight := getGifDimensions(p.gif)

	p.rgba = image.NewRGBA(image.Rect(0, 0, gifWidth, gifHeight))
	if p.rgba.Stride != p.rgba.Rect.Size().X*4 {
		return errors.New("unsupported stride\n")
	}

	draw.Draw(p.rgba, p.rgba.Bounds(), p.gif.Image[0], image.ZP, draw.Over)

	return nil
}

// InitPluginTexture is used to create and manage backend textures
func (p *GifPlugin) InitPluginTexture(registry *flutter.TextureRegistry) error {

	texture := registry.NewTexture()

	// for this exmaple I have hard-coded a `textureId: 2` in the flutter
	// Texture widget.
	texture.ID = 2
	// You should never change this ID, `registry.NewTexture()` will keep track
	// of the textureId counter for you.
	// You should send the texture.ID number to dart using some sort of platform
	// channel, and use the go generated textureId in flutter.

	go func() {
		for { // looping gif
			for i, srcImg := range p.gif.Image {
				time.Sleep(time.Millisecond * 10 * time.Duration(p.gif.Delay[i]))
				draw.Draw(p.rgba, p.rgba.Bounds(), srcImg, image.ZP, draw.Over) // get the frame's pixels
				texture.FrameAvailable()                                        // redraw (calls textureHanler, and retrive the pixel)
			}
		}
		// In this fairly simple example no golang mutex/channel were used to sync draw.Draw and textureHanler.
	}()

	return texture.Register(p.textureHanler)
}

// textureHanler is executed on the main thread, try to make this handle as
// light as possible.
// In this case the handler is only sending the pixel buffer drawn in a
// goroutine.
func (p *GifPlugin) textureHanler(width, height int) (bool, *flutter.PixelBuffer) {
	return true, &flutter.PixelBuffer{
		Pix:    p.rgba.Pix,
		Width:  p.rgba.Bounds().Size().X,
		Height: p.rgba.Bounds().Size().Y,
	}
}

func getGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX int
	var lowestY int
	var highestX int
	var highestY int
	for _, gif := range gif.Image {
		if gif.Rect.Min.X < lowestX {
			lowestX = gif.Rect.Min.X
		}
		if gif.Rect.Min.Y < lowestY {
			lowestY = gif.Rect.Min.Y
		}
		if gif.Rect.Max.X > highestX {
			highestX = gif.Rect.Max.X
		}
		if gif.Rect.Max.Y > highestY {
			highestY = gif.Rect.Max.Y
		}
	}
	return highestX - lowestX, highestY - lowestY
}
