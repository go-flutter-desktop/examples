import 'package:flutter/foundation.dart'
    show debugDefaultTargetPlatformOverride;

import 'package:flutter/material.dart';
import 'package:flame/flame.dart';
import 'package:flutter/gestures.dart';

import './game.dart';

void main() async {
  debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;
  final Size size = await Flame.util.initialDimensions();
  final game = MyGame(size);
  runApp(game.widget);

  Flame.util.addGestureRecognizer(TapGestureRecognizer()
    ..onTapDown = (TapDownDetails evt) {
      game.searchController();
    });
}
