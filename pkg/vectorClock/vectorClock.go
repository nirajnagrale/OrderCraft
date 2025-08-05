package vectorClock

import (
	"maps"
	"sync"
)

// VectorClock represents a vector clock for distributed systems.

type VectorClock struct {
	mu    sync.Mutex
	clock map[string]int
	id    string
}

type VcSnap struct {
	Clock map[string]int
	Id    string
}

// NewVectorClock initializes a new vector clock with the given size.
func NewVectorClock(id string) *VectorClock {
	vc := &VectorClock{
		clock: make(map[string]int),
		id:    id,
	}
	vc.clock[id] = 0
	return vc
}

// GetSnapshot returns a snapshot of the current vector clock.
func (vc *VectorClock) GetSnapshot() VcSnap {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	snapshot := VcSnap{
		Clock: make(map[string]int),
		Id:    vc.id,
	}
	maps.Copy(snapshot.Clock, vc.clock)
	return snapshot
}

// event on every vectorClock
func (vc *VectorClock) Tick() (VcSnap, error) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	vc.clock[vc.id]++

	snapshot := VcSnap{
		Clock: make(map[string]int),
		Id:    vc.id,
	}
	maps.Copy(snapshot.Clock, vc.clock)
	return snapshot, nil
}

// Add some node to vc
func (vc *VectorClock) AddNode(id string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	_, exists := vc.clock[id]
	if !exists {
		vc.clock[id] = 0
	}
}

// delete some node if exists
func (vc *VectorClock) DeleteNode(id string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	_, exists := vc.clock[id]
	if exists {
		delete(vc.clock, id)
	}
}
