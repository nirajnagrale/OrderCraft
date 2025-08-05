package fifo

import (
	"fmt"
	"ordercraft/pkg/vectorClock"
)

// vc1 is owner clock
// vc2 is the clock of the node we want to compare with
// true = can be delivered, 1 = cannot be delivered
// returns 0 if vc1 can be delivered, 1 if it cannot be delivered
func CompareClocks(vc1, vc2 vectorClock.VcSnap) (bool, error) {
	id := vc2.Id
	_, exists := vc1.Clock[id]
	if !exists {
		return false, fmt.Errorf("id %s not found in vc1", id)
	}
	if vc1.Clock[id]+1 == vc2.Clock[id] {
		return true, nil
	}
	return false, nil
}
