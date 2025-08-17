package messageQueue_test

import (
	messageQueue "ordercraft/pkg/MessageQueue"
	vectorClock "ordercraft/pkg/VectorClock"
	"testing"
)

func TestMessageQueue(t *testing.T) {
	// Create a new message queue
	mq := messageQueue.NewMessageQueue()
	// Create vector clocks for testing
	vc1 := vectorClock.NewVectorClock("node1")
	vc2 := vectorClock.NewVectorClock("node2")

	// Create queue nodes with messages
	vc1_snapshot := vc1.GetSnapshot()
	vc2_snapshot := vc2.GetSnapshot()

	node1 := messageQueue.GetQueueItem(&vc1_snapshot, "Message from node1")
	node2 := messageQueue.GetQueueItem(&vc2_snapshot, "Message from node2")

	// Enqueue messages
	mq.Buffer(node1)
	mq.Buffer(node2)
	// Check the size of the queue for node1
	if len(mq.Nodes["node1"]) != 1 {
		t.Errorf("Expected 1 message in node1's queue, got %d", len(mq.Nodes["node1"]))
	}
	// Check the size of the queue for node2
	if len(mq.Nodes["node2"]) != 1 {
		t.Errorf("Expected 1 message in node2's queue, got %d", len(mq.Nodes["node2"]))
	}
	// Dequeue messages
	dequeuedNode1, exists1 := mq.Deliver("node1")
	if !exists1 || dequeuedNode1.Msg != "Message from node1" {
		t.Errorf("Expected to dequeue 'Message from node1', got %s", dequeuedNode1.Msg)
	}
	dequeuedNode2, exists2 := mq.Deliver("node2")
	if !exists2 || dequeuedNode2.Msg != "Message from node2" {
		t.Errorf("Expected to dequeue 'Message from node2', got %s", dequeuedNode2.Msg)
	}
	// Check if the queues are empty after dequeueing
	if len(mq.Nodes["node1"]) != 0 {
		t.Errorf("Expected node1's queue to be empty, got %d", len(mq.Nodes["node1"]))
	}

}
