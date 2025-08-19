package messageBuffer_test

import (
	messageBuffer "ordercraft/pkg/MessageBuffer"
	vectorClock "ordercraft/pkg/VectorClock"
	payload "ordercraft/pkg/Payload"
	"testing"
)

func TestMessageBuffer(t *testing.T) {
	// Create a new message buffer
	mq := messageBuffer.NewBuffer()
	// Create vector clocks for testing
	vc1 := vectorClock.NewVectorClock("node1")
	vc2 := vectorClock.NewVectorClock("node2")

	// Create queue nodes with messages
	vc1_snapshot := vc1.GetSnapshot()
	vc2_snapshot := vc2.GetSnapshot()

	payload1 := payload.MakePayload(&vc1_snapshot, "Message from node1")
	payload2 := payload.MakePayload(&vc2_snapshot, "Message from node2")

	// Enqueue messages
	mq.Buffer(payload1)
	mq.Buffer(payload2)
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
