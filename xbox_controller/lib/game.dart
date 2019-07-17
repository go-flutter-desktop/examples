import 'package:flame/game.dart';
import 'package:flame/svg.dart';
import 'package:flame/position.dart';
import 'package:flame/text_config.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'dart:async';

const XBOXEventStream = const EventChannel('go-flutter-sample/xbox_controller');

class MyGame extends BaseGame {
  Svg svgInstance;
  TextConfig config = TextConfig(
      fontSize: 20.0, fontFamily: 'Roboto', color: Colors.pinkAccent);

  Controller controller = Controller.messageAsName(
      "No game controller connected (Click to connect)");

  StreamSubscription _xboxInputSubs;

  Size screenSize;
  MyGame(this.screenSize) {
    _start();
    searchController();
  }

  void searchController() {
    if (_xboxInputSubs != null) {
      return;
    }
    this.controller = Controller.messageAsName("Searching controller...");

    _xboxInputSubs =
        XBOXEventStream.receiveBroadcastStream().listen(_updateGameController);
    _xboxInputSubs.onDone(() async {
      await Future.delayed(Duration(seconds: 1));
      this.controller = Controller.messageAsName(
          "No game controller connected (Click to connect)");
      _xboxInputSubs = null;
    });
  }

  void _updateGameController(ctrl) {
    // print(ctrl);
    this.controller = Controller(
        name: ctrl["name"],
        stickAxes: ctrl["axes"].cast<double>(),
        buttons: ctrl["buttons"].cast<bool>());
  }

  void _start() {
    svgInstance = Svg('xbox-controller.svg');
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    final Rect rect = Rect.fromLTWH(0, 0, screenSize.width, screenSize.height);

    canvas.drawRect(rect, Paint()..color = Colors.white);
    canvas.drawRect(rect, Paint()..color = Colors.black12);

    config.render(canvas, this.controller.name, Position(10, 10));

    svgInstance.renderPosition(canvas, Position(100, 0), 600, 500);

    canvas.drawCircle(
        Offset(269 + this.controller.offsetRightStick().x,
            174 + this.controller.offsetRightStick().y),
        7,
        Paint()..color = Colors.white70);

    canvas.drawCircle(
        Offset(466 + this.controller.offsetLeftStick().x,
            260 + this.controller.offsetLeftStick().y),
        7,
        Paint()..color = Colors.white70);

    canvas.drawRect(Rect.fromLTWH(259, 91, 19, this.controller.lt() * 10),
        Paint()..color = Colors.black54);

    canvas.drawRect(Rect.fromLTWH(523, 91, 19, this.controller.rt() * 10),
        Paint()..color = Colors.black54);

    for (var btn in this.controller.pressedButtons()) {
      canvas.drawCircle(Offset(btn.coordinates.x, btn.coordinates.y), 18,
          Paint()..color = btn.color);
    }
  }
}

class Coordinates {
  double x, y;
  Coordinates({this.x, this.y});
}

class Button {
  String name;
  Coordinates coordinates;
  Color color;
  Button(this.name, this.coordinates, this.color);
}

class Controller {
  String name;
  List<double> stickAxes = [0, 0, 0, 0, 0, 0, 0, 0];
  List<bool> buttons = [
    false,
    false,
    false,
    false,
    false,
    false,
    false,
    false,
    false,
    false
  ];

  Controller({this.name, this.stickAxes, this.buttons});

  // Named constructor
  Controller.messageAsName(this.name);

  Coordinates offsetRightStick() {
    return Coordinates(x: (this.stickAxes[0] * 27), y: (stickAxes[1] * 27));
  }

  Coordinates offsetLeftStick() {
    return Coordinates(x: (this.stickAxes[2] * 27), y: (stickAxes[3] * 27));
  }

  double rt() {
    return (this.stickAxes[4] + 1) * -2;
  }

  double lt() {
    return (this.stickAxes[5] + 1) * -2;
  }

  List<Button> pressedButtons() {
    var list = List<Button>();
    if (this.buttons[0])
      list.add(Button("A", Coordinates(x: 526, y: 208), Color(0xFF4AB072)));
    if (this.buttons[1])
      list.add(Button("B", Coordinates(x: 561, y: 174), Color(0xFFBC322B)));
    if (this.buttons[2])
      list.add(Button("X", Coordinates(x: 492, y: 174), Color(0xFF355BB8)));
    if (this.buttons[3])
      list.add(Button("Y", Coordinates(x: 526, y: 140), Color(0xFFCFAF2B)));

    if (this.stickAxes[6] == 1)
      list.add(Button(">", Coordinates(x: 384, y: 257), Colors.white70));
    if (this.stickAxes[6] == -1)
      list.add(Button("<", Coordinates(x: 292, y: 257), Colors.white70));
    if (this.stickAxes[7] == -1)
      list.add(Button("^", Coordinates(x: 339, y: 210), Colors.white70));
    if (this.stickAxes[7] == 1)
      list.add(Button("v", Coordinates(x: 339, y: 307), Colors.white70));

    if (this.buttons[5])
      list.add(Button("RB", Coordinates(x: 535, y: 83), Colors.white70));
    if (this.buttons[4])
      list.add(Button("LB", Coordinates(x: 266, y: 83), Colors.white70));
    return list;
  }
}
