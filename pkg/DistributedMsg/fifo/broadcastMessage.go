package fifo

import (
	"fmt"
	"math/rand"
	"net"
	messageQueue "ordercraft/pkg/MessageQueue"
	nodesip "ordercraft/pkg/Nodesip"
	vectorClock "ordercraft/pkg/VectorClock"
	"sync"
	"time"
)

func broadcastMessages(vc *vectorClock.VectorClock, nodesIp *nodesip.NodesIp,
	noOfMessages int) {
	// Implementation of broadcasting messages
	// This function would typically send messages to other nodes
	// in the distributed system, updating the vector clock accordingly.
	messages := make([]string, noOfMessages)
	for i := range noOfMessages {
		messages[i] = "Message " + (string)(i)
	}
	connections := make([]net.Conn, 0, len(*nodesIp))
	var wg sync.WaitGroup
	wg.Add(len(*nodesIp) * len(messages)) // Add the number of nodes to wait for
	defer wg.Wait()
	for _, ip := range *nodesIp {
		conn, err := net.DialTimeout("tcp", ip, 20*time.Second)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Connected to node at", ip)
		}
		connections = append(connections, conn)
	}
	for i := range messages {
		vcSnapshot, err := vc.Tick()
		if err != nil {
			fmt.Println("Error ticking vector clock:", err)
			continue
		}
		queueItem := messageQueue.GetQueueItem(&vcSnapshot, messages[i])	
		for _, conn := range connections {
			go handleBroadcast(queueItem,conn)
		}
	}
}

func handleBroadcast(queueItem messageQueue.QueueItem, conn net.Conn) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate network delay
	defer conn.Close()
	if err != nil {
		fmt.Println("Error marshalling queue item:", err)
		return
	}
	_, err = conn.Write([]byte(queueItem))
}
