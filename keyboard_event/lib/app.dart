import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter/services.dart';

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Code Sample for Keyboard event',
      theme: ThemeData(
        // If the host is missing some fonts, it can cause the
        // text to not be rendered or worse the app might crash.
        fontFamily: 'Roboto',
        primarySwatch: Colors.blue,
      ),
      home: KeyboardTestPage(),
    );
  }
}

/// Keyboard test page for the example application.
class KeyboardTestPage extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    return _KeyboardTestPageState();
  }
}

class _KeyboardTestPageState extends State<KeyboardTestPage> {
  final List<String> _messages = [];

  final FocusNode _focusNode = FocusNode();
  final ScrollController _scrollController = new ScrollController();

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    FocusScope.of(context).requestFocus(_focusNode);
  }

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      appBar: AppBar(
        title: new Text('Keyboard events test'),
      ),
      body: new RawKeyboardListener(
        focusNode: _focusNode,
        onKey: onKeyEvent,
        child: Container(
          padding: EdgeInsets.symmetric(horizontal: 16.0),
          child: SingleChildScrollView(
              controller: _scrollController,
              child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: _messages.map((m) => new Text(m)).toList())),
        ),
      ),
    );
  }

  void onKeyEvent(RawKeyEvent event) {
    bool isKeyDown;
    switch (event.runtimeType) {
      case RawKeyDownEvent:
        isKeyDown = true;
        break;
      case RawKeyUpEvent:
        isKeyDown = false;
        break;
      default:
        throw new Exception('Unexpected runtimeType of RawKeyEvent');
    }

    int keyCode;
    switch (event.data.runtimeType) {
      case RawKeyEventDataLinux:
        final RawKeyEventDataLinux data = event.data;
        keyCode = data.keyCode;
        break;
      default:
        throw new Exception('Unsupported platform ${event.data.runtimeType}');
    }

    _addMessage(
        '${isKeyDown ? 'KeyDown' : 'KeyUp'}: $keyCode\n- Modifiers: ${event.data.modifiersPressed}\n- KeyLabel: ${event.data.keyLabel}\n- physicalKey: ${event.data.physicalKey}\n- character: ${event.character}');
  }

  void _addMessage(String message) {
    setState(() {
      _messages.add(message);
    });
    SchedulerBinding.instance.addPostFrameCallback((_) {
      _scrollController.jumpTo(
        _scrollController.position.maxScrollExtent,
      );
    });
  }
}
