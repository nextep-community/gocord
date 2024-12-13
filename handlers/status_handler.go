package handlers

import (
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func gatewayHandlerRaw(client bot.Client, sequenceNumber int, shardID int, event gateway.EventRaw) {
	client.EventManager().DispatchEvent(&events.Raw{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		EventRaw:     event,
	})
}

func gatewayHandlerHeartbeatAck(client bot.Client, sequenceNumber int, shardID int, event gateway.EventHeartbeatAck) {
	client.EventManager().DispatchEvent(&events.HeartbeatAck{
		GenericEvent:      events.NewGenericEvent(client, sequenceNumber, shardID),
		EventHeartbeatAck: event,
	})
}

func gatewayHandlerReady(client bot.Client, sequenceNumber int, shardID int, event gateway.EventReady) {
	client.Caches().SetSelfUser(event.User)

	for _, guild := range event.Guilds {
		client.Caches().SetGuildUnready(guild.ID, true)
	}

	client.EventManager().DispatchEvent(&events.Ready{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		EventReady:   event,
	})
}

func gatewayHandlerResumed(client bot.Client, sequenceNumber int, shardID int, _ gateway.EventData) {
	client.EventManager().DispatchEvent(&events.Resumed{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
	})
}
