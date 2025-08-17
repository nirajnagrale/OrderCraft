package fifo

import (
	"fmt"
	messageQueue "ordercraft/pkg/MessageQueue"
	nodesip "ordercraft/pkg/Nodesip"
	vectorClock "ordercraft/pkg/VectorClock"
	"sync"
)

// vc1 is owner clock
// vc2 is the clock of the node we want to compare with
// true = can be delivered, 1 = cannot be delivered
// returns 0 if vc1 can be delivered, 1 if it cannot be delivered
func compareClocks(vc1, vc2 *vectorClock.VcSnap) bool {
	id := vc2.Id
	_, exists := vc1.Clock[id]
	if !exists {
		return false
	}
	if vc1.Clock[id]+1 == vc2.Clock[id] {
		return true
	}
	return false
}

func CanDeliverMsg(vc1, vc2 *vectorClock.VcSnap) bool {
	canDeliver := compareClocks(vc1, vc2)

	if canDeliver {
		// Update the clock to reflect the delivery
		vc1.Clock[vc1.Id] = vc2.Clock[vc2.Id]
		return true
	}
	return false
}

func Run(vc *vectorClock.VectorClock, mq *messageQueue.MessageQueue, nodesIp *nodesip.NodesIp, noOfMessages int, address string) {
	// Implementation of the Run function
	// This function would typically handle the message queue processing
	// and the vector clock updates based on the messages received.
	// receive messages and broadcast 3 random messages
	fmt.Println("Running FIFO with address:", address)
	var wg sync.WaitGroup
	wg.Add(2) // Two goroutines: one for broadcasting and one for receiving
	go broadcastMessages(vc, nodesIp, noOfMessages)
	go receiveMessages(vc, mq, address)
	wg.Wait()

}
