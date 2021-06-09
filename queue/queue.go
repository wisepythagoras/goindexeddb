package queue

// Queue defines our queue structure.
type Queue struct {
	itemsInQueue []*Job
}

// Enqueue adds a job to the queue.
func (q *Queue) Enqueue(job *Job) {
	q.itemsInQueue = append(q.itemsInQueue, job)
}

// Dequeue returns the oldest job and removes it from the queue.
func (q *Queue) Dequeue() *Job {
	if len(q.itemsInQueue) == 0 {
		return nil
	}

	job := q.itemsInQueue[0]
	q.itemsInQueue = q.itemsInQueue[1:]

	return job
}
