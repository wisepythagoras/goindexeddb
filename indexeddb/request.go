package indexeddb

type ReadyStateType string

const (
	RequestStatePending ReadyStateType = "pending"
	RequestStateDone    ReadyStateType = "done"
	RequestStateError   ReadyStateType = "error"
)

// DBInternalRequest implements the core database request.
type DBInternalRequest struct {
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
	*DBInternalRequest
}
