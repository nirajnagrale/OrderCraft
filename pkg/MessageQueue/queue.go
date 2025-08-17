package messageQueue

import vectorClock "ordercraft/pkg/VectorClock"

// MessageQueue represents a message queue for distributed systems.
type QueueItem struct {
	Vc  *vectorClock.VcSnap // Vector clock for the node (pointer)
	Msg string              // message from the id
}

type queue []QueueItem

func GetQueueItem(vc *vectorClock.VcSnap, msg string) QueueItem {
	return QueueItem{
		Vc:  vc,
		Msg: msg,
	}
}

func MakeQueue() []QueueItem {
	return make([]QueueItem, 0)
}
