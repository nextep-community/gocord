package handlers

import (
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func gatewayHandlerWebhooksUpdate(client bot.Client, sequenceNumber int, shardID int, event gateway.EventWebhooksUpdate) {
	client.EventManager().DispatchEvent(&events.WebhooksUpdate{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		GuildId:      event.GuildID,
		ChannelID:    event.ChannelID,
	})
}
