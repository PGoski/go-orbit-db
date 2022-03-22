package pubsub

import (
	"github.com/libp2p/go-eventbus"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/peer"

	"berty.tech/go-orbit-db/iface"
)

type Event interface{}

type PayloadEmitter struct {
	event.Emitter
}

func NewPayloadEmitter(bus event.Bus) (*PayloadEmitter, error) {
	emitter, err := bus.Emitter(new(iface.EventPubSubPayload), eventbus.Stateful)
	if err != nil {
		return nil, err
	}

	return &PayloadEmitter{emitter}, nil

}

func (e *PayloadEmitter) Emit(evt *iface.EventPubSubPayload) error {
	return e.Emitter.Emit(*evt)
}

// Creates a new Message event
func NewEventMessage(content []byte, p peer.ID) *iface.EventPubSubMessage {
	return &iface.EventPubSubMessage{
		Content: content,
		From:    p,
	}
}

// NewEventPayload Creates a new Message event
func NewEventPayload(payload []byte, p peer.ID) *iface.EventPubSubPayload {
	return &iface.EventPubSubPayload{
		Payload: payload,
		From:    p,
	}
}

// NewEventPeerJoin creates a new EventPubSubJoin event
func NewEventPeerJoin(p peer.ID) Event {
	return &iface.EventPubSubJoin{
		Peer: p,
	}
}

// NewEventPeerLeave creates a new EventPubSubLeave event
func NewEventPeerLeave(p peer.ID) Event {
	return &iface.EventPubSubLeave{
		Peer: p,
	}
}
