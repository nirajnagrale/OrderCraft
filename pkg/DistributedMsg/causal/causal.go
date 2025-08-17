package causal

import (
	"ordercraft/pkg/VectorClock"
)

func compareClocks(vc1, vc2 *vectorClock.VcSnap) bool {
	id := vc2.Id
	_, exists := vc1.Clock[id]
	if !exists {
		return false
	}
	for k, v := range vc1.Clock {
		if k == id {
			if v+1 != vc2.Clock[k] {
				return false // vc1 is not less than vc2
			}
			continue // skip the owner clock comparison
		}
		if v < vc2.Clock[k] {
			return false // vc1 is not less than vc2
		}
	}
	return true
}

func canDeliverMsg(vc1, vc2 *vectorClock.VcSnap) bool {
	canDeliver := compareClocks(vc1, vc2)

	if canDeliver {
		// Update the clock to reflect the delivery
		vc1.Clock[vc1.Id] = vc2.Clock[vc2.Id]
		return true
	}
	return false
}
