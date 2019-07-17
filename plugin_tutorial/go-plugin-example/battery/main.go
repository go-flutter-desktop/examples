package battery

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

//  Make sure to use the same channel name as was used on the Flutter client side.
const channelName = "samples.flutter.dev/battery"

// MyBatteryPlugin demonstrates how to call a platform-specific API to retrieve
// the current battery level.
//
// This example matches the guide example available on:
// https://flutter.dev/docs/development/platform-integration/platform-channels
type MyBatteryPlugin struct{}

var _ flutter.Plugin = &MyBatteryPlugin{} // compile-time type check

// InitPlugin creates a MethodChannel and set a HandleFunc to the
// shared 'getBatteryLevel' method.
// https://godoc.org/github.com/go-flutter-desktop/go-flutter/plugin#MethodChannel
//
// The plugin is using a MethodChannel through the StandardMethodCodec.
//
// You can also use the more basic BasicMessageChannel, which supports basic,
// asynchronous message passing using a custom message codec.
// You can also use the specialized BinaryCodec, StringCodec, and JSONMessageCodec
// struct, or create your own codec.
//
// The JSONMessageCodec was deliberately left out because in Go its better to
// decode directly to known structs.
func (p *MyBatteryPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getBatteryLevel", handleGetBatteryLevel)
	return nil // no error
}

// handleGetBatteryLevel is called when the method getBatteryLevel is invoked by
// the dart code.
// The function returns a fake battery level, without raising an error.
//
// Supported return types of StandardMethodCodec codec are described in a table:
// https://godoc.org/github.com/go-flutter-desktop/go-flutter/plugin#StandardMessageCodec
func handleGetBatteryLevel(arguments interface{}) (reply interface{}, err error) {
	batteryLevel := int32(55) // Your platform-specific API
	return batteryLevel, nil
}
