package messageQueue

import (
	"sync"
)

type MessageQueue struct {
	mu    sync.Mutex
	Nodes map[string]queue // Map of nodes with their vector clocks and messages
}

// NewMessageQueue creates a new MessageQueue instance.
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		Nodes: make(map[string]queue, 0),
	}
}

func (mq *MessageQueue) Enqueue(incoming_node QueueItem) {
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
	q = append(q, QueueItem{})
	copy(q[insertPos+1:], q[insertPos:])
	q[insertPos] = incoming_node
	mq.Nodes[senderID] = q
}

func (mq *MessageQueue) Dequeue(senderID string) (QueueItem, bool) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	q, exists := mq.Nodes[senderID]
	if !exists || len(q) == 0 {
		return QueueItem{}, false // No messages to dequeue
	}

	// Dequeue the first message from the queue
	dequeuedNode := q[0]
	q = q[1:] // Remove the first element
	mq.Nodes[senderID] = q
	return dequeuedNode, true
}
