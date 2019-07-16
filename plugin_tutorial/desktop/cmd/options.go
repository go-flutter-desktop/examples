package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	battery "github.com/go-flutter-desktop/plugins/go-plugin-example/battery"
	"github.com/go-flutter-desktop/plugins/go-plugin-example/complex_data_structure"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(100, 100),
	flutter.PopBehavior(flutter.PopBehaviorClose),        // on SystemNavigator.pop() closes the app
	flutter.AddPlugin(&battery.MyBatteryPlugin{}),        // our wiki plugin
	flutter.AddPlugin(&complex_data_structure.Example{}), // another example
}
