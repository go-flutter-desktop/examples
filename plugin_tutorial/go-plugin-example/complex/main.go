package complex

import (
	"errors"
	"log"
	"os"
	"time"

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
	p.channel.HandleFunc("mutualCall", p.mutualCall)
	p.channel.HandleFunc("getError", getErrorFunc)
	p.channel.CatchAllHandleFunc(catchAllTest)

	return nil
}

// mutualCall is attached to the plugin struct
func (p *Example) mutualCall(arguments interface{}) (reply interface{}, err error) {
	go func() {
		time.Sleep(3 * time.Second)
		if rep, err := p.channel.InvokeMethodWithReply("InvokeMethodWithReply", "text_from_golang"); err != nil {
			log.Println("InvokeMethod error:", err)
		} else {
			if rep.(string) != "text_from_dart" {
				log.Println("InvokeMethod error: rep.(string) != \"text_from_dart\"")
				os.Exit(1)
			}
		}
	}()

	if arguments.(string) != "hello_from_dart" {
		log.Println("InvokeMethod error: arguments.(string) !=  \"hello_from_dart\"")
		os.Exit(1)
	}

	return "hello_from_go", nil
}

func catchAllTest(methodCall interface{}) (reply interface{}, err error) {
	method := methodCall.(plugin.MethodCall)
	// return the randomized Method Name
	return method.Method, nil
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

func getErrorFunc(arguments interface{}) (reply interface{}, err error) {
	return nil, plugin.NewError("customErrorCode", errors.New("Some error"))
}
