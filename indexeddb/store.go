package indexeddb

type DBObjectStore struct {
	DB *Database
}

func (o *DBObjectStore) Init() {
	// Check if object store exists exists.
}

func (o *DBObjectStore) Add(object *interface{}) {}
