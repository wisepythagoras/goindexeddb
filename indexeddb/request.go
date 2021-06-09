package indexeddb

type ReadyStateType string

const (
	RequestStatePending ReadyStateType = "pending"
	RequestStateDone    ReadyStateType = "done"
	RequestStateError   ReadyStateType = "error"
)

// InternalRequest implements the core database request.
type InternalRequest struct {
	Source      interface{}
	Transaction interface{}
	ReadyState  ReadyStateType
	Result      *Database
	OnSuccess   *CallbackFn
	OnError     *CallbackFn
}

// DBRequest implements the database request.
type DBRequest struct {
	*EventTarget
	*InternalRequest
}
