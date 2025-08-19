package messageBuffer

import (
	payload "ordercraft/pkg/Payload"
)

// MessageQueue represents a message queue for distributed systems.
type queue []payload.Payload

func MakeQueue() []payload.Payload {
	return make([]payload.Payload, 0)
}
