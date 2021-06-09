package indexeddb

import (
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
	Path   string
	dbName string
}

// Init initializes the DBFactory.
func (f *DBFactory) Init() {
	fmt.Println(f.Path)
}

// Open opens the DBFactory.
func (f *DBFactory) Open(name string, version int) (*DBOpenDBRequest, *sync.WaitGroup, error) {
	eventTarget := &EventTarget{}
	internalRequest := &DBInternalRequest{
		ReadyState: RequestStatePending,
	}
	request := &DBRequest{eventTarget, internalRequest}
	openRequest := &DBOpenDBRequest{request}

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
			connectionError := CreateEvent("error", "Unable to connect")

			if openRequest.OnError != nil {
				(*openRequest.OnError)(connectionError)
			}

			openRequest.DispatchEvent(connectionError)

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

		openRequest.Result = database
		openRequest.ReadyState = RequestStateDone
		openRequest.DispatchEvent(CreateEvent("upgradeneeded", false))
		openRequest.DispatchEvent(CreateEvent("success", true))
	}()

	return openRequest, &wg, nil
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
