package messageBuffer

import (
	payload "ordercraft/pkg/Payload"
	"sync"
)

type Buffer struct {
	mu    sync.Mutex
	Nodes map[string]queue // Map of nodes with their vector clocks and messages
}

// NewBuffer creates a new Buffer instance.
func NewBuffer() *Buffer {
	return &Buffer{
		Nodes: make(map[string]queue, 0),
	}
}

func (mq *Buffer) Buffer(incoming_node payload.Payload) {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	//check the position of the incoming node
	// keep it in asceding order of vector clock
	senderID := incoming_node.Vc.Id
	q := mq.Nodes[senderID]
	insertPos := 0
	for i, n := range q {
		if n.Vc.Clock[senderID] < incoming_node.Vc.Clock[senderID] {
			insertPos = i + 1
		} else {
			break
		}
	}
	// Insert the incoming node at the correct position
	q = append(q, payload.Payload{})
	copy(q[insertPos+1:], q[insertPos:])
	q[insertPos] = incoming_node
	mq.Nodes[senderID] = q
}

func (mq *Buffer) Deliver(senderID string) (payload.Payload, bool) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	q, exists := mq.Nodes[senderID]
	if !exists || len(q) == 0 {
		return payload.Payload{}, false // No messages to dequeue
	}

	// Dequeue the first message from the queue
	dequeuedNode := q[0]
	q = q[1:] // Remove the first element
	mq.Nodes[senderID] = q
	return dequeuedNode, true
}
