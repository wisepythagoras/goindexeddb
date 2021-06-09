package indexeddb

type InternalOpenDBRequest struct {
	OnOpen  *CallbackFn
	OnClose *CallbackFn
}

// DBOpenDBRequest implements the request for opening the DB.
type DBOpenDBRequest struct {
	*DBRequest
	*InternalOpenDBRequest
}
