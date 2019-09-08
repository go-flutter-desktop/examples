package example_image

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// ImagePlugin is a texture plugin example
type ImagePlugin struct{}

var _ flutter.PluginTexture = &ImagePlugin{} // compile-time type check

// InitPlugin is used because PluginTexture must implement flutter.Plugin
func (p *ImagePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	return nil
}

// InitPluginTexture is used to create and manage backend textures
func (p *ImagePlugin) InitPluginTexture(registry *flutter.TextureRegistry) error {

	texture := registry.NewTexture()

	// for this exmaple I hard-coded a `textureId: 1` int the flutter Texture widget.
	texture.ID = 1
	// You should never change this ID, `registry.NewTexture()` will keep track
	// of the textureId counter for you.
	// You should send the texture.ID number to dart using some sort of platform
	// channel, you should use the go generated textureId in flutter.

	// after 5 seconds, remove the texture from the scene.
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Printf("texture_tutorial: UnRegistering texture %v after 5 seconds\n", texture.ID)
		err := texture.UnRegister()
		if err != nil {
			fmt.Printf("texture_tutorial: %v\n", err)
		}
		texture.FrameAvailable() // redraw
	}()

	return texture.Register(p.textureHanler)
}

func (p *ImagePlugin) textureHanler(width, height int) (bool, *flutter.PixelBuffer) {

	// hard-coded path, run the app with `hover run` in the project root
	file := "./test.png"
	imgFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("texture %q not found on disk: %v\n", file, err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Printf("error decoding file %s:%v\n", file, err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		fmt.Printf("unsupported stride\n")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return true, &flutter.PixelBuffer{
		Pix:    rgba.Pix,
		Width:  img.Bounds().Size().X,
		Height: img.Bounds().Size().Y,
	}
}
