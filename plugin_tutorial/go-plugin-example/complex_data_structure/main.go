package complex_data_structure

import (
	"errors"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// Example demonstrates how to call a platform-specific API to retrieve
// a complex data structure
type Example struct{}

var _ flutter.Plugin = &Example{}

// InitPlugin creates a MethodChannel and set a HandleFunc to the
// shared 'getData' method.
func (Example) InitPlugin(messenger plugin.BinaryMessenger) error {

	channel := plugin.NewMethodChannel(messenger, "instance.id/go/data", plugin.StandardMethodCodec{})
	channel.HandleFunc("getData", getRemotesFunc)

	return nil
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
