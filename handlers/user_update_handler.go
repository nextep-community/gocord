package handlers

import (
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func gatewayHandlerUserUpdate(client bot.Client, sequenceNumber int, shardID int, event gateway.EventUserUpdate) {
	oldUser, _ := client.Caches().SelfUser()
	client.Caches().SetSelfUser(event.OAuth2User)

	client.EventManager().DispatchEvent(&events.SelfUpdate{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber, shardID),
		SelfUser:     event.OAuth2User,
		OldSelfUser:  oldUser,
	})

}
