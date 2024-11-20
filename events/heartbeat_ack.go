package events

import "github.com/nextep-community/gocord/gateway"

type HeartbeatAck struct {
	*GenericEvent
	gateway.EventHeartbeatAck
}
