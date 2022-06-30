// This example makes a [Container] react to being entered by a mouse
// pointer, showing a count of the number of entries and exits.

import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Code Sample for widgets.Listener',
      theme: ThemeData(
        // If the host is missing some fonts, it can cause the
        // text to not be rendered or worse the app might crash.
        fontFamily: 'Roboto',
        primarySwatch: Colors.blue,
      ),
      home: MyStatefulWidget(),
    );
  }
}

class MyStatefulWidget extends StatefulWidget {
  MyStatefulWidget({Key key}) : super(key: key);

  @override
  _MyStatefulWidgetState createState() => _MyStatefulWidgetState();
}

class _MyStatefulWidgetState extends State<MyStatefulWidget> {
  int _enterCounter = 0;
  int _exitCounter = 0;
  double x = 0.0;
  double y = 0.0;
  String _mouseStatus = "hovering";

  void _incrementCounter(PointerEnterEvent details) {
    setState(() {
      _enterCounter++;
    });
  }

  void _decrementCounter(PointerExitEvent details) {
    setState(() {
      _exitCounter++;
    });
  }

  void _updateLocation(PointerHoverEvent details) {
    setState(() {
      x = details.position.dx;
      y = details.position.dy;
      _mouseStatus = "hovering, button ${details.buttons.toString()}";
    });
  }

  void _updateMove(PointerMoveEvent details) {
    setState(() {
      x = details.position.dx;
      y = details.position.dy;
      _mouseStatus = "dragging, button ${details.buttons.toString()}";
    });
  }

  void _pointerDown(PointerDownEvent details) {
    setState(() {
      x = details.position.dx;
      y = details.position.dy;
      _mouseStatus = "down, button ${details.buttons.toString()}";
    });
  }

  void _pointerUp(PointerUpEvent details) {
    setState(() {
      x = details.position.dx;
      y = details.position.dy;
      _mouseStatus = "up, button ${details.buttons.toString()}";
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Hover Example'),
      ),
      body: Center(
        child: ConstrainedBox(
          constraints: BoxConstraints.tight(Size(300.0, 200.0)),
          child: MouseRegion(
            onEnter: _incrementCounter,
            onExit: _decrementCounter,
            child: Listener(
              onPointerDown: _pointerDown,
              onPointerUp: _pointerUp,
              onPointerMove: _updateMove,
              onPointerHover: _updateLocation,
              child: Container(
                color: Colors.lightBlueAccent,
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Text('You have pointed at this box this many times:'),
                    Text(
                      '$_enterCounter Entries\n$_exitCounter Exits',
                    ),
                    Text(
                      'The cursor is here: (${x.toStringAsFixed(2)}, ${y.toStringAsFixed(2)})',
                    ),
                    Divider(),
                    Text(
                      'The mouse is ${_mouseStatus}',
                    ),
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
