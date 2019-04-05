import 'package:flutter/foundation.dart' show debugDefaultTargetPlatformOverride;
import 'package:flutter/material.dart';

import 'package:simple_demo/app.dart';

void main() {
  // Desktop platforms aren't a valid platform.
  debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;
  runApp(new MyApp());
}
