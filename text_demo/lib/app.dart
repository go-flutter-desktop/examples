// This example demos the TextField/SelectableText widget and keyboard
// integration with the go-flutter text backend

import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Code Sample for testing text input',
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
  FocusNode myFocus = FocusNode();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Hover Example'),
      ),
      body: Container(
        child: Column(
          children: [
            Padding(
              padding: EdgeInsets.all(16.0),
              child: SelectableText.rich(
                // this text can be copied with "Ctrl-C"
                TextSpan(
                  text: 'Hello', // default text style
                  children: <TextSpan>[
                    TextSpan(
                        text: ' beautiful world',
                        style: TextStyle(fontWeight: FontWeight.bold)),
                    TextSpan(
                        text: ' (should be copyable)',
                        style: TextStyle(
                            fontStyle: FontStyle.italic, fontSize: 10)),
                  ],
                ),
              ),
            ),
            Padding(
              padding: new EdgeInsets.all(8.0),
              child: new Column(children: <Widget>[
                TextField(
                  obscureText: false,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(),
                    labelText: 'TextField',
                  ),
                  onSubmitted: (value) {
                    print("TextField 1:" + value);
                  },
                  onEditingComplete: () =>
                      FocusScope.of(context).requestFocus(myFocus),
                ),
                TextField(
                  decoration: InputDecoration(hintText: 'Multiline TextField'),
                  maxLines: 4,
                  focusNode: myFocus,
                  onSubmitted: (value) {
                    setState(() {
                      print("TextField 2:" + value);
                    });
                  },
                ),
              ]),
            ),
          ],
        ),
      ),
    );
  }
}
