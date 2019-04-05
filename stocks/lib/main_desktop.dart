import 'package:flutter/foundation.dart';
import 'package:flutter/widgets.dart';
import 'package:stocks/stock.dart';

void main() {
  debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;
  runApp(new StocksApp());
}
