package indexeddb

import (
	"fmt"
	"sync"
)

// CreateEvent creates a dispatchable event.
func CreateEvent(eventType string, bubbles interface{}) *Event {
	return &Event{
		Type:    eventType,
		Bubbles: bubbles,
	}
}

// DBFactory implements the main DB instance.
type DBFactory struct {
	Path string
}

// Init initializes the DBFactory.
func (f *DBFactory) Init() {
	fmt.Println(f.Path)
}

// Open opens the DBFactory.
func (f *DBFactory) Open(name string, version int) (*DBOpenDBRequest, *sync.WaitGroup, error) {
	eventTarget := &EventTarget{}
	internalRequest := &DBInternalRequest{}
	request := &DBRequest{eventTarget, internalRequest}
	openRequest := &DBOpenDBRequest{request}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// Do stuff here.
		for i := 0; i < 1000000000; i++ {
		}

		openRequest.DispatchEvent(CreateEvent("upgradeneeded", false))
		openRequest.DispatchEvent(CreateEvent("success", true))
	}()

	return openRequest, &wg, nil
}
