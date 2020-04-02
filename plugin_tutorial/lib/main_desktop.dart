import 'package:flutter/services.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_test/flutter_test.dart';
import 'dart:math';
import 'dart:async' as async;

// The client and host sides of a channel are connected through
// a channel name passed in the channel constructor.
const platform_channel_battery =
    const MethodChannel('samples.flutter.dev/battery');

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
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

  group('Complex plugin', () {
    const platform_complex_structure =
        const MethodChannel('instance.id/go/data');

    // A more complicated plugin
    test('Test StandardMethodCodec array of map', () async {
      final List<dynamic> result =
          await platform_complex_structure.invokeMethod('getData',
              "HelloFromDart"); // passing "HelloFromDart" as an argument
      expect(result, [
        {"instanceid": 1023, "pcbackup": "test", "brbackup": "test2"},
        {"instanceid": 1024, "pcbackup": "test", "brbackup": "test2"},
        {"instanceid": 1056, "pcbackup": "coucou", "brbackup": "coucou2"},
        {"instanceid": 3322, "pcbackup": "finaly", "brbackup": "finaly2"}
      ]);
    });

    test('Test golang InvokeMethodWithReply', () async {
      var method;
      var arguments;
      async.Completer completer = new async.Completer();

      platform_complex_structure.setMethodCallHandler((MethodCall call) async {
        method = call.method;
        arguments = call.arguments;
        completer.complete();
        return "text_from_dart";
      });

      // Triggers the above setMethodCallHandler handler
      var result = await platform_complex_structure.invokeMethod(
          'mutualCall', 'hello_from_dart');
      expect(result, "hello_from_go");

      await completer.future;

      expect(method, "InvokeMethodWithReply");
      expect(arguments, "text_from_golang");
    });

    test('Custom errors', () async {
      var matcher = predicate(
          (e) => e is PlatformException && e.code == "customErrorCode");
      expect(
        platform_complex_structure.invokeMethod('getError'),
        throwsA(matcher),
      );
    });

    test('Test golang CatchAllHandleFunc', () async {
      // golang can return the random methodName
      final String methodName =
          'test/' + new Random().nextInt(100000).toString();
      final String resultPathPrefix =
          await platform_complex_structure.invokeMethod(methodName);
      expect(resultPathPrefix, methodName);
    });
  });

  tearDownAll(() async {
    SystemNavigator.pop(); // close the app
  });
}
