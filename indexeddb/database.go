package indexeddb

type KeyOptions struct {
	AutoIncrement bool
	KeyPath       string
}

type Database struct {
	Factory *DBFactory
	Name    string
	Version int
}

// CreateObjectStore creates a new object store.
func (db *Database) CreateObjectStore(name string, keyOpts *KeyOptions) *DBObjectStore {
	objectStore := &DBObjectStore{
		DB: db,
	}

	return objectStore
}
