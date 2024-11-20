package handlers

import (
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func gatewayHandlerGuildBanAdd(client bot.Client, sequenceNumber int, shardID int, event gateway.EventGuildBanAdd) {
	client.EventManager().DispatchEvent(&events.GuildBan{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		GuildID:      event.GuildID,
		User:         event.User,
	})
}

func gatewayHandlerGuildBanRemove(client bot.Client, sequenceNumber int, shardID int, event gateway.EventGuildBanRemove) {
	client.EventManager().DispatchEvent(&events.GuildUnban{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		GuildID:      event.GuildID,
		User:         event.User,
	})
}
