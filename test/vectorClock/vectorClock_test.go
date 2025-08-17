package test

import (
	vectorClock "ordercraft/pkg/VectorClock"
	"testing"
)

func TestVectorClock(t *testing.T) {
	// Initialize a new vector clock
	vc := vectorClock.NewVectorClock("node1")

	// Tick the vector clock
	snapshot, err := vc.Tick()
	if err != nil {
		panic(err)
	}

	// Print the snapshot id
	println("Snapshot after tick:")
	println("Snapshot ID:", snapshot.Id)
	//print the snapshot values
	for id, value := range snapshot.Clock {
		println("Node:", id, "Value:", value)
	}
	// Add a new node
	vc.AddNode("node2")
	println("Added node2 to vector clock")

	// Tick again
	snapshot, err = vc.Tick()
	if err != nil {
		panic(err)
	}
	println("Snapshot after adding node2:")
	println("Snapshot after tick:", snapshot.Id)
	//print the snapshot values
	for id, value := range snapshot.Clock {
		println("Node:", id, "Value:", value)
	}
	// Delete a node
	vc.DeleteNode("node1")
	println("Deleted node1 from vector clock")
}
