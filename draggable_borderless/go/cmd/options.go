package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(400, 300),
	flutter.WindowMode(flutter.WindowModeBorderless),
	flutter.AddPlugin(&AppBarDraggable{}),
}

// AppBarDraggable is a plugin that makes moving the bordreless window possible
type AppBarDraggable struct {
	window           *glfw.Window
	windowDragActive chan bool
}

var _ flutter.Plugin = &AppBarDraggable{}     // compile-time type check
var _ flutter.PluginGLFW = &AppBarDraggable{} // compile-time type check
// AppBarDraggable struct must implement InitPlugin and InitPluginGLFW

// InitPlugin creates a MethodChannel for "samples.go-flutter.dev/draggable"
func (p *AppBarDraggable) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.windowDragActive = make(chan bool)
	channel := plugin.NewMethodChannel(messenger, "samples.go-flutter.dev/draggable", plugin.StandardMethodCodec{})
	channel.HandleFunc("onPanStart", p.onPanStart)
	channel.HandleFunc("onPanEnd", p.onPanEnd)
	return nil
}

// InitPluginGLFW is used to gain control over the glfw.Window
func (p *AppBarDraggable) InitPluginGLFW(window *glfw.Window) error {
	p.window = window
	return nil
}

// onPanStart a golang / flutter implemantation of:
// "GLFW how to drag undecorated window without lag"
// https://stackoverflow.com/a/46205940
func (p *AppBarDraggable) onPanStart(arguments interface{}) (reply interface{}, err error) {
	argumentsMap := arguments.(map[interface{}]interface{})
	cursorPosX := int(argumentsMap["dx"].(float64))
	cursorPosY := int(argumentsMap["dy"].(float64))
	for {
		select {
		case <-p.windowDragActive:
			return
		default:
			xpos, ypos := p.window.GetCursorPos()
			deltaX := int(xpos) - cursorPosX
			deltaY := int(ypos) - cursorPosY

			x, y := p.window.GetPos()
			p.window.SetPos(x+deltaX, y+deltaY)
		}

	}
	return nil, nil
}

func (p *AppBarDraggable) onPanEnd(arguments interface{}) (reply interface{}, err error) {
	p.windowDragActive <- false
	return nil, nil
}
