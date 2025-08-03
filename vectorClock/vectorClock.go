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
	clock map[string]int
	id    string
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

//event on every vectorClock
func (vc *VectorClock) Tick() (VcSnap, error) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.clock[vc.id]++
	var vc_snap VcSnap
	maps.Copy(vc_snap.clock, (vc.clock))
	vc_snap.id = vc.id
	return vc_snap, nil
}

//Add some node to vc
func (vc *VectorClock) AddNode(id string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	_, exists := vc.clock[id]
	if !exists {
		vc.clock[id] = 0
	}
}
//delete some node if exists
func (vc *VectorClock) DeleteNode(id string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	_, exists := vc.clock[id]; if exists {
		delete(vc.clock,id)
	}
}

