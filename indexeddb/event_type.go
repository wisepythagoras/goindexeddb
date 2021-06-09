package indexeddb

type EventType string

const (
	EventTypeOpen          EventType = "open"
	EventTypeSuccess       EventType = "success"
	EventTypeUpgradeNeeded EventType = "upgradeneeded"
	EventTypeError         EventType = "error"
)
