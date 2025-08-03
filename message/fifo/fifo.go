package fifo

import (
	"ordercraft/vectorClock"
	"sync"
)

// FifoQueue represents a FIFO queue for distributed systems.
type FifoQueue struct {
	mu    sync.Mutex
	queue []vectorClock.VectorClock
	id    string // Unique identifier for the queue
	msg   string // message from the id 
}

// NewFifoQueue initializes a new FIFO queue.
func NewFifoQueue() *FifoQueue {
	return &FifoQueue{
		queue: make([]vectorClock.VectorClock, 0),
	}
}

// Enqueue adds an item to the end of the queue.
func (f *FifoQueue) Enqueue(item vectorClock.VectorClock) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.queue = append(f.queue, item)
}

// Dequeue removes and returns the item at the front of the queue.
func (f *FifoQueue) Dequeue() (string, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.queue) == 0 {
		return "", false // Queue is empty
	}
	item := f.queue[0]
	f.queue = f.queue[1:] // Remove the first item
	return item, true
}
