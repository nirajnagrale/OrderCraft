package fifo

import (
	"encoding/json"
	"fmt"
	"net"
	messageBuffer "ordercraft/pkg/MessageBuffer"
	payload "ordercraft/pkg/Payload"
	vectorClock "ordercraft/pkg/VectorClock"
)

func receiveMessages(vc *vectorClock.VectorClock, mq *messageBuffer.Buffer, address string) {
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

func handleIncomingMessage(vc *vectorClock.VectorClock, mq *messageBuffer.Buffer, conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 1024) // Adjust buffer size as needed
	n, err := conn.Read(data)
	if err != nil {
		return // Handle error appropriately
	}
	// Deserialize the message (assuming it's a string for simplicity)
	var Payload payload.Payload
	err = json.Unmarshal(data[:n], &Payload)
	if err != nil {
		fmt.Println("Error unmarshalling message:", err)
		return // Handle error appropriately
	}

	vcSnapshot := vc.GetSnapshot()

	// Buffer the message in the message queue
	if CanDeliverMsg(&vcSnapshot, Payload.Vc) {
		fmt.Println("Message from node", Payload.Vc.Id, "can be delivered with message:", Payload.Msg)
	} else {
		mq.Buffer(Payload)
	}
}

func checkMessagQueue(vc *vectorClock.VectorClock, mq *messageBuffer.Buffer) {
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
