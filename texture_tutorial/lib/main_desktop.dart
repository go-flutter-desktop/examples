import 'package:flutter/foundation.dart'
    show debugDefaultTargetPlatformOverride;
import 'package:flutter/material.dart';
import 'package:video_player/video_player.dart';

void main() {
  debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'go-flutter Texture API',
      theme: ThemeData(
        // If the host is missing some fonts, it can cause the
        // text to not be rendered or worse the app might crash.
        fontFamily: 'Roboto',
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        appBar: AppBar(
          title: Text('Texture Examples'),
        ),
        body: Center(
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: <Widget>[
              Column(
                children: <Widget>[
                  Spacer(flex: 2),
                  Text('Video Player'),
                  Spacer(flex: 1),
                  ConstrainedBox(
                    constraints: BoxConstraints.tight(Size(640, 360)),
                    // showcase the video_player plugin
                    child: VideoPlayer(VideoPlayerController.network(
                        'SampleVideo_1280x720_10mb.mp4')
                      ..initialize()
                      ..play()),
                  ),
                  Spacer(flex: 2),
                ],
              ),
              Column(
                children: <Widget>[
                  Spacer(flex: 2),
                  Text('Image Texture'),
                  Spacer(flex: 1),
                  ConstrainedBox(
                    constraints: BoxConstraints.tight(Size(330, 319)),
                    child: Texture(textureId: 1), // hard-coded to 1, image
                  ),
                  Spacer(flex: 2),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
