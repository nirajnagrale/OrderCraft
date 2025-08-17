package fifo_test

import (
	"ordercraft/pkg/DistributedMsg/fifo"
	messageQueue "ordercraft/pkg/MessageQueue"
	nodesip "ordercraft/pkg/Nodesip"
	vectorClock "ordercraft/pkg/VectorClock"
	"testing"
)

func TestFifo(t *testing.T) {
	vc1 := vectorClock.NewVectorClock("node1")
	vc2 := vectorClock.NewVectorClock("node2")
	//vc3 := vectorClock.NewVectorClock("node3")

	nodesIp := nodesip.NewNodesIp()
	nodesIp.AddNode("node1", "localhost:8081")
	nodesIp.AddNode("node2", "localhost:8082")
	//nodesIp.AddNode("node3", "localhost:8083")

	mq1 := messageQueue.NewMessageQueue()
	mq2 := messageQueue.NewMessageQueue()
	//mq3 := messageQueue.NewMessageQueue()

	go fifo.Run(vc1, mq1, nodesIp, 1, "localhost:8081")
	go fifo.Run(vc2, mq2, nodesIp, 1, "localhost:8082")
	//go fifo.Run(vc3, mq3, nodesIp, 1, "localhost:8083")

	select {}
}
