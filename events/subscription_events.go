package events

import "github.com/nextep-community/gocord/discord"

type GenericSubscriptionEvent struct {
	*GenericEvent
	discord.Subscription
}

type SubscriptionCreate struct {
	*GenericSubscriptionEvent
}

type SubscriptionUpdate struct {
	*GenericSubscriptionEvent
}

type SubscriptionDelete struct {
	*GenericSubscriptionEvent
}
