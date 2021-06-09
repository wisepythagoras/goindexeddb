package indexeddb

type ReadyStateType string

const (
	Pending ReadyStateType = "pending"
)

// DBInternalRequest implements the core database request.
type DBInternalRequest struct {
	Source      interface{}
	Transaction interface{}
	ReadyState  ReadyStateType
	OnSuccess   *CallbackFn
	OnError     *CallbackFn
}

// DBRequest implements the database request.
type DBRequest struct {
	*EventTarget
	*DBInternalRequest
}
