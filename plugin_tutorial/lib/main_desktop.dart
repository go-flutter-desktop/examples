import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'dart:math';

// The client and host sides of a channel are connected through
// a channel name passed in the channel constructor.
const platform_channel_battery =
    const MethodChannel('samples.flutter.dev/battery');

void main() async {
  // Example from the WIKI
  //
  test('Test Battery Plugin result', () async {
    // Invoke a method on the method channel, specifying the concrete
    // method to call via the String identifier.
    final int result =
        await platform_channel_battery.invokeMethod('getBatteryLevel');
    expect(result, 55);

    final String batteryLevel = '          -- Battery level at $result % .';
    print(batteryLevel);
  });

  // A more complicated plugin
  //
  test('Test StandardMethodCodec array of map', () async {
    const platform_complex_structure =
        const MethodChannel('instance.id/go/data');

    final List<dynamic> result = await platform_complex_structure.invokeMethod(
        'getData', "HelloFromDart"); // passing "HelloFromDart" as an argument
    expect(result, [
      {"instanceid": 1023, "pcbackup": "test", "brbackup": "test2"},
      {"instanceid": 1024, "pcbackup": "test", "brbackup": "test2"},
      {"instanceid": 1056, "pcbackup": "coucou", "brbackup": "coucou2"},
      {"instanceid": 3322, "pcbackup": "finaly", "brbackup": "finaly2"}
    ]);

    // golang can return the random methodName
    final String methodName = 'test/' + new Random().nextInt(100000).toString();
    final String resultPathPrefix =
        await platform_complex_structure.invokeMethod(methodName);
    expect(resultPathPrefix, methodName);
  });

  tearDownAll(() async {
    SystemNavigator.pop(); // close the app
  });
}
