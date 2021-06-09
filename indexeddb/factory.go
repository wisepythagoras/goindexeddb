package indexeddb

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	Path        string
	dbName      string
	openRequest *DBOpenDBRequest
}

// Init initializes the DBFactory.
func (f *DBFactory) Init() {
	fmt.Println(f.Path)
}

// Open opens the DBFactory.
func (f *DBFactory) Open(name string, version int) (*DBOpenDBRequest, *sync.WaitGroup, error) {
	eventTarget := &EventTarget{}
	internalRequest := &InternalRequest{
		ReadyState: RequestStatePending,
	}
	request := &DBRequest{eventTarget, internalRequest}
	f.openRequest = &DBOpenDBRequest{request, &InternalOpenDBRequest{}}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		path := f.Path

		if len(path) > 0 {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				err := os.MkdirAll(path, os.ModeDir|os.ModePerm)

				if err != nil {
					panic(err)
				}
			}

			if path[len(path)-1] != '/' {
				path = path + "/"
			}
		}

		// Try to open the database.
		db, err := gorm.Open(sqlite.Open(path+name+".db"), &gorm.Config{})

		if err != nil {
			openError := errors.New("Unable to connect")
			err = f.dispatchEvent(EventTypeError, openError)

			if err != nil {
				panic(err)
			}

			return
		}

		db.AutoMigrate(&IDBTableDataKey{})
		db.AutoMigrate(&IDBTableKey{})
		db.AutoMigrate(&IDBTableData{})
		db.AutoMigrate(&IDBTable{})
		db.AutoMigrate(&IDBDatabase{})

		database := &Database{
			Factory: f,
			Name:    name,
			Version: version,
		}

		f.dbName = name

		f.openRequest.Result = database
		f.openRequest.ReadyState = RequestStateDone
		f.dispatchEvent(EventTypeOpen, f.openRequest.Result)
		f.dispatchEvent(EventTypeSuccess, f.openRequest.Result)
	}()

	return f.openRequest, &wg, nil
}

// Delete removes the database from the filesystem.
func (f *DBFactory) Delete() error {
	path := f.Path

	if len(path) > 0 && path[len(path)-1] != '/' {
		path = path + "/"
	}

	_, err := os.Stat(path)

	if err != nil {
		return os.Remove(path + f.dbName + "db")
	}

	return err
}

// dispatchEvent creates and dispatches an event.
func (f *DBFactory) dispatchEvent(t EventType, payload interface{}) error {
	if f.openRequest == nil {
		return errors.New("Database is not open")
	}

	eventData := CreateEvent(string(t), payload)

	if t == EventTypeOpen {
		if f.openRequest.OnOpen != nil {
			(*f.openRequest.OnOpen)(eventData)
		}
	} else if t == EventTypeError {
		if f.openRequest.OnError != nil {
			(*f.openRequest.OnError)(eventData)
		}
	} else if t == EventTypeSuccess {
		if f.openRequest.OnSuccess != nil {
			(*f.openRequest.OnSuccess)(eventData)
		}
	} else {
		return errors.New("Invalid event type")
	}

	f.openRequest.DispatchEvent(eventData)

	return nil
}
