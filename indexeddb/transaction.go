package indexeddb

type InternalTransaction struct {
	OnComplete *CallbackFn
}

type DBTransaction struct {
	*EventTarget
	*InternalTransaction
}
