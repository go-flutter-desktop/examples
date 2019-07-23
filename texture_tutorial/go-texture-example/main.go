package example

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// MyTexturePlugin is a texture plugin example
type MyTexturePlugin struct{}

var _ flutter.PluginTexture = &MyTexturePlugin{} // compile-time type check

// InitPlugin is used because PluginTexture must implement flutter.Plugin
func (p *MyTexturePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	return nil
}

// InitPluginTexture is used to create and manage backend textures
func (p *MyTexturePlugin) InitPluginTexture(registry *flutter.TextureRegistry) error {

	texture := registry.NewTexture()

	// for this exmaple I hardcoded a `textureId: 1` for the flutter Texture widget.
	texture.ID = 1

	// `registry.NewTexture()` will keep track of the textureId counter.
	// You should send the texture.ID number to dart using some sort of platform
	// channel. Don't handle the texture.ID by yourself!!, it's an example.

	// after 20 seconds, remove the texture from the scene.
	go func() {
		time.Sleep(20 * time.Second)
		fmt.Printf("UnRegistering texture %v\n", texture.ID)
		err := texture.UnRegister()
		if err != nil {
			fmt.Printf("example.MyTexturePlugin: %v", err)
		}
		texture.FrameAvailable() // repaint now!
	}()

	return texture.Register(p.textureHanler)
}

func (p *MyTexturePlugin) textureHanler(width, height int) (bool, *flutter.PixelBuffer) {

	// hard coded the app with hover run
	file := "./test.png"
	imgFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("texture %q not found on disk: %v", file, err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Printf("error decoding file %s:%v", file, err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		fmt.Printf("unsupported stride")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return true, &flutter.PixelBuffer{
		Pix:    rgba.Pix,
		Width:  img.Bounds().Size().X,
		Height: img.Bounds().Size().Y,
	}
}
