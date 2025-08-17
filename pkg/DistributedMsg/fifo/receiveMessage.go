package fifo

import (
	"encoding/json"
	"fmt"
	"net"
	messageQueue "ordercraft/pkg/MessageQueue"
	vectorClock "ordercraft/pkg/VectorClock"
)

func receiveMessages(vc *vectorClock.VectorClock, mq *messageQueue.MessageQueue, address string) {
	go checkMessagQueue(vc, mq)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue // Handle error appropriately
		}
		go handleIncomingMessage(vc, mq, conn)
	}
}

func handleIncomingMessage(vc *vectorClock.VectorClock, mq *messageQueue.MessageQueue, conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024) // Adjust buffer size as needed
	n, err := conn.Read(buffer)
	if err != nil {
		return // Handle error appropriately
	}
	// Deserialize the message (assuming it's a string for simplicity)
	message := string(buffer[:n])

	// Create a new vector clock snapshot
	vcSnapshot := vc.GetSnapshot()
	b, _ := json.Marshal(vcSnapshot.Clock)
	fmt.Printf("Received message from ID=%s, VC=%s\n", vcSnapshot.Id, string(b))

	// Create a queue item with the received message and vector clock
	queueItem := messageQueue.QueueItem{
		Vc:  &vcSnapshot,
		Msg: message,
	}

	// Buffer the message in the message queue
	if CanDeliverMsg(&vcSnapshot, queueItem.Vc) {
		fmt.Println("Message from node", queueItem.Vc.Id, "can be delivered with message:", queueItem.Msg)
	} else {
		mq.Buffer(queueItem)
	}
}

func checkMessagQueue(vc *vectorClock.VectorClock, mq *messageQueue.MessageQueue) {
	for {
		for id, queue := range mq.Nodes {
			if len(queue) > 0 {
				queueItem := queue[0] // Get the first item in the queue
				vc_snap := vc.GetSnapshot()
				if CanDeliverMsg(&vc_snap, queueItem.Vc) {
					fmt.Println("Delivering message from node", queueItem.Vc.Id, "with message:", queueItem.Msg)
					mq.Deliver(id) // Deliver the message
				}
			}
		}
	}
}
