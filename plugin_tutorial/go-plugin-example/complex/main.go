package complex

import (
	"errors"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// Example demonstrates how to call a platform-specific API to retrieve
// a complex data structure
type Example struct {
	channel *plugin.MethodChannel
}

var _ flutter.Plugin = &Example{}

// InitPlugin creates a MethodChannel and set a HandleFunc to the
// shared 'getData' method.
func (p *Example) InitPlugin(messenger plugin.BinaryMessenger) error {

	p.channel = plugin.NewMethodChannel(messenger, "instance.id/go/data", plugin.StandardMethodCodec{})
	p.channel.HandleFunc("getData", getRemotesFunc)

	// The order of PathPrefix is important!
	p.channel.PathPrefix("test/").HandleFunc(p.pathPrefixTest) // "test/" will match before ""
	p.channel.PathPrefix("").HandleFunc(matchAll)

	return nil
}

func (p *Example) pathPrefixTest(methodCall interface{}) (reply interface{}, err error) {
	method := methodCall.(plugin.MethodCall)

	// // delete HandleFunc on "test/"
	// next call to "test/" will use the MethodCall "matchAll"
	p.channel.PathPrefix("test/").HandleFunc(nil)

	return method.Method, nil
}

func matchAll(methodCall interface{}) (reply interface{}, err error) {
	return "matchAll", nil
}

func getRemotesFunc(arguments interface{}) (reply interface{}, err error) {
	dartMsg := arguments.(string) // reading the string argument
	if dartMsg != "HelloFromDart" {
		return nil, errors.New("wrong message value, expecting 'HelloFromDart' got '" + dartMsg + "'")
	}

	var sectionList = make([]interface{}, 4)
	sectionList[0] = map[interface{}]interface{}{
		"instanceid": int32(1023),
		"pcbackup":   "test",
		"brbackup":   "test2",
	}
	sectionList[1] = map[interface{}]interface{}{
		"instanceid": int32(1024),
		"pcbackup":   "test",
		"brbackup":   "test2",
	}

	sectionList[2] = map[interface{}]interface{}{
		"instanceid": int32(1056),
		"pcbackup":   "coucou",
		"brbackup":   "coucou2",
	}

	sectionList[3] = map[interface{}]interface{}{
		"instanceid": int32(3322),
		"pcbackup":   "finaly",
		"brbackup":   "finaly2",
	}

	return sectionList, nil
}
