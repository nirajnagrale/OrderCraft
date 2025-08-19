package payload

import vectorClock "ordercraft/pkg/VectorClock"


type Payload struct {
	Vc  *vectorClock.VcSnap // Vector clock for the node (pointer)
	Msg string              // message from the id
}

func MakePayload(vc *vectorClock.VcSnap, msg string) Payload {
	return Payload{
		Vc:  vc,
		Msg: msg,
	}
}
