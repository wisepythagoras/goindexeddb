package queue

// Job describes a queue job
type Job struct {
	JobType string
	Data    interface{}
}
