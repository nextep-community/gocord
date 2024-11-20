package handlers

import (
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func gatewayHandlerInviteCreate(client bot.Client, sequenceNumber int, shardID int, event gateway.EventInviteCreate) {
	client.EventManager().DispatchEvent(&events.InviteCreate{
		GenericEvent:      events.NewGenericEvent(client, sequenceNumber, shardID),
		EventInviteCreate: event,
	})
}

func gatewayHandlerInviteDelete(client bot.Client, sequenceNumber int, shardID int, event gateway.EventInviteDelete) {
	client.EventManager().DispatchEvent(&events.InviteDelete{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		GuildID:      event.GuildID,
		ChannelID:    event.ChannelID,
		Code:         event.Code,
	})
}
