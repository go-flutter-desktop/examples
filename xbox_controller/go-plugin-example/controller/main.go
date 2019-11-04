package controller

import (
	"reflect"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// XBOXStream demonstrates how to use the EventChannel.
type XBOXStream struct {
	stop chan bool
}

var _ flutter.Plugin = &XBOXStream{} // compile-time type check

func (p *XBOXStream) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.stop = make(chan bool)
	channel := plugin.NewEventChannel(messenger, "go-flutter-sample/xbox_controller", plugin.StandardMethodCodec{})
	channel.Handle(p)
	return nil
}

func (p *XBOXStream) OnListen(arguments interface{}, sink *plugin.EventSink) {
	var lastJoystickInfo map[interface{}]interface{}
	for {
		select {
		case <-p.stop:
			return
		default:

			var present bool
			joy := glfw.Joystick1
			if glfw.JoystickPresent(joy) {
				present = true

				glfwButtons := glfw.GetJoystickButtons(joy)
				buttons := make([]interface{}, len(glfwButtons))
				for i, v := range glfwButtons {
					buttons[i] = v == 1
				}

				glfwAxes := glfw.GetJoystickAxes(joy)
				axes := make([]interface{}, len(glfwAxes))
				for i, v := range glfwAxes {
					axes[i] = float64(v)
				}

				joystickInfo := map[interface{}]interface{}{
					"name":    glfw.GetJoystickName(joy),
					"buttons": buttons,
					"axes":    axes,
				}

				// don't send duplicated succeding message, (reduces serialisation).
				if !reflect.DeepEqual(lastJoystickInfo, joystickInfo) {
					sink.Success(joystickInfo)
					lastJoystickInfo = joystickInfo
				}
			}
			if !present {
				// no joysticks present

				// EndOfStream will call OnCancel, which will send true in the channel
				// p.stop, which will stop the OnListen goroutine.
				sink.EndOfStream()
			}

			time.Sleep(5 * time.Millisecond) // Constantly polling joysticks event
			// no wait for joysticks events available
		}
	}
}
func (p *XBOXStream) OnCancel(arguments interface{}) {
	// I choose to use channels to Cancel events.
	// Mutex can also work.
	// I found that channels are bit more reliable than Mutex during hot restart.
	p.stop <- true
}
