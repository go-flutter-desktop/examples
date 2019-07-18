package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/plugins/go-plugin-example/battery"
	"github.com/go-flutter-desktop/plugins/go-plugin-example/complex"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(100, 100),
	flutter.PopBehavior(flutter.PopBehaviorClose), // on SystemNavigator.pop() closes the app
	flutter.AddPlugin(&battery.MyBatteryPlugin{}), // our wiki plugin
	flutter.AddPlugin(&complex.Example{}),         // another example
}
