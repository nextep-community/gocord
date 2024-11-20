package handlers

import (
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func gatewayHandlerApplicationCommandPermissionsUpdate(client bot.Client, sequenceNumber int, shardID int, event gateway.EventApplicationCommandPermissionsUpdate) {
	client.EventManager().DispatchEvent(&events.GuildApplicationCommandPermissionsUpdate{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		Permissions:  event.ApplicationCommandPermissions,
	})
}
