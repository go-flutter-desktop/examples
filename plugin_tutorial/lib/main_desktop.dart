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

    // invokeMethod on a wildcard MethodHandler
    final String matchAll = new Random().nextInt(100000).toString();
    final String resultPathPrefixMatchall =
        await platform_complex_structure.invokeMethod(matchAll);
    expect(resultPathPrefixMatchall, "matchAll");

    // The golang `channel.PathPrefix("test/")` was added BEFORE the wildcard
    // handler, no conflict should occur with `channel.PathPrefix("")`
    //
    // The golang MethodHandler "test/" will delete itself.
    final String methodName = 'test/' + new Random().nextInt(100000).toString();
    final String resultPathPrefix =
        await platform_complex_structure.invokeMethod(methodName);
    expect(resultPathPrefix, methodName);

    // Another call on "test/" should use the wildcard MethodHandler
    final String resultPathPrefix2 =
        await platform_complex_structure.invokeMethod(methodName);
    expect(resultPathPrefix2, "matchAll");
  });

  tearDownAll(() async {
    SystemNavigator.pop(); // close the app
  });
}
