# draggable_borderless

A [go-flutter-desktop](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info) GLFW plugin example.  
In this example we are answering issue [#214](https://github.com/go-flutter-desktop/go-flutter/issues/214): *"how to drag undecorated window".*  

The GLFW answer comes from https://stackoverflow.com/a/46205940 with a little
twist.  
The window is only draggable from the [Flutter AppBar](https://api.flutter.dev/flutter/material/AppBar-class.html) widget.

The source code of the plugin has been placed directly in the [`desktop/cmd/options.go`](./desktop/cmd/options.go) file.  
A custom `AppBar` widget has been created, the source code is available in the
[`lib/main_desktop.dart`](./lib/main_desktop.dart) file.

## Demo
<p align="center">
  <img src="./app_bar_drag.gif" width="650" align="center" alt="Demo of the
  example">
</p>
