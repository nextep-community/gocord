package events

import "github.com/nextep-community/gocord/gateway"

type Raw struct {
	*GenericEvent
	gateway.EventRaw
}
