package handlers

import (
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func gatewayHandlerGuildRoleCreate(client bot.Client, sequenceNumber int, shardID int, event gateway.EventGuildRoleCreate) {
	client.Caches().AddRole(event.Role)

	client.EventManager().DispatchEvent(&events.RoleCreate{
		GenericRole: &events.GenericRole{
			GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
			GuildID:      event.GuildID,
			RoleID:       event.Role.ID,
			Role:         event.Role,
		},
	})
}

func gatewayHandlerGuildRoleUpdate(client bot.Client, sequenceNumber int, shardID int, event gateway.EventGuildRoleUpdate) {
	oldRole, _ := client.Caches().Role(event.GuildID, event.Role.ID)
	client.Caches().AddRole(event.Role)

	client.EventManager().DispatchEvent(&events.RoleUpdate{
		GenericRole: &events.GenericRole{
			GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
			GuildID:      event.GuildID,
			RoleID:       event.Role.ID,
			Role:         event.Role,
		},
		OldRole: oldRole,
	})
}

func gatewayHandlerGuildRoleDelete(client bot.Client, sequenceNumber int, shardID int, event gateway.EventGuildRoleDelete) {
	role, _ := client.Caches().RemoveRole(event.GuildID, event.RoleID)

	client.EventManager().DispatchEvent(&events.RoleDelete{
		GenericRole: &events.GenericRole{
			GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
			GuildID:      event.GuildID,
			RoleID:       event.RoleID,
			Role:         role,
		},
	})
}
