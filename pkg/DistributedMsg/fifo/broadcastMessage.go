package fifo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	nodesip "ordercraft/pkg/Nodesip"
	payload "ordercraft/pkg/Payload"
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
		Payload := payload.MakePayload(&vcSnapshot, messages[i])
		for _, conn := range connections {
			go handleBroadcast(Payload, conn)
		}
	}
}

func handleBroadcast(Payload payload.Payload, conn net.Conn) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate network delay
	data, err := json.Marshal(Payload)
	defer conn.Close()
	if err != nil {
		fmt.Println("Error marshalling queue item:", err)
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}
}
