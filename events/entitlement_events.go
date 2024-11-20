package events

import "github.com/nextep-community/gocord/discord"

type GenericEntitlementEvent struct {
	*GenericEvent
	discord.Entitlement
}

type EntitlementCreate struct {
	*GenericEntitlementEvent
}

type EntitlementUpdate struct {
	*GenericEntitlementEvent
}

type EntitlementDelete struct {
	*GenericEntitlementEvent
}
