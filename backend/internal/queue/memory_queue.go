package queue

import "sync"

type MemoryQueue struct {
	mu   sync.Mutex
	jobs []Payload
}

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{jobs: make([]Payload, 0)}
}

func (q *MemoryQueue) Enqueue(payload Payload) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.jobs = append(q.jobs, payload)
	return nil
}

func (q *MemoryQueue) Jobs() []Payload {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.jobs
}
