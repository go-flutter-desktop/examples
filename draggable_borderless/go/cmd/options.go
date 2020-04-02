package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(400, 300),
	flutter.WindowMode(flutter.WindowModeBorderless),
	flutter.AddPlugin(&AppBarDraggable{}),
}

// AppBarDraggable is a plugin that makes moving the bordreless window possible
type AppBarDraggable struct {
	window           *glfw.Window
	cursorPosY       int
	cursorPosX       int
}

var _ flutter.Plugin = &AppBarDraggable{}     // compile-time type check
var _ flutter.PluginGLFW = &AppBarDraggable{} // compile-time type check
// AppBarDraggable struct must implement InitPlugin and InitPluginGLFW

// InitPlugin creates a MethodChannel for "samples.go-flutter.dev/draggable"
func (p *AppBarDraggable) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, "samples.go-flutter.dev/draggable", plugin.StandardMethodCodec{})
	channel.HandleFunc("onPanStart", p.onPanStart)
	channel.HandleFuncSync("onPanUpdate", p.onPanUpdate) // MUST RUN ON THE MAIN THREAD (use of HandleFuncSync)
	channel.HandleFunc("onClose", p.onClose)
	return nil
}

// InitPluginGLFW is used to gain control over the glfw.Window
func (p *AppBarDraggable) InitPluginGLFW(window *glfw.Window) error {
	p.window = window
	return nil
}

// onPanStart/onPanUpdate a golang / flutter implemantation of:
// "GLFW how to drag undecorated window without lag"
// https://stackoverflow.com/a/46205940
func (p *AppBarDraggable) onPanStart(arguments interface{}) (reply interface{}, err error) {
	argumentsMap := arguments.(map[interface{}]interface{})
	p.cursorPosX = int(argumentsMap["dx"].(float64))
	p.cursorPosY = int(argumentsMap["dy"].(float64))
	return nil, nil
}

// onPanUpdate calls GLFW functions that aren't thread safe.
// to run function on the main go-flutter thread, use HandleFuncSync instead of HandleFunc!
func (p *AppBarDraggable) onPanUpdate(arguments interface{}) (reply interface{}, err error) {
	xpos, ypos := p.window.GetCursorPos() // This function must only be called from the main thread.
	deltaX := int(xpos) - p.cursorPosX
	deltaY := int(ypos) - p.cursorPosY

	x, y := p.window.GetPos()           // This function must only be called from the main thread.
	p.window.SetPos(x+deltaX, y+deltaY) // This function must only be called from the main thread.

	return nil, nil
}

func (p *AppBarDraggable) onClose(arguments interface{}) (reply interface{}, err error) {
	// This function may be called from any thread. Access is not synchronized.
	p.window.SetShouldClose(true)
	return nil, nil
}
