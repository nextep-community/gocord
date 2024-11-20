package handlers

import (
	"log/slog"

	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/discord"
	"github.com/nextep-community/gocord/httpserver"
)

var _ bot.HTTPServerEventHandler = (*httpserverHandlerInteractionCreate)(nil)

type httpserverHandlerInteractionCreate struct{}

func (h *httpserverHandlerInteractionCreate) HandleHTTPEvent(client bot.Client, respondFunc httpserver.RespondFunc, event httpserver.EventInteractionCreate) {
	// we just want to pong all pings
	// no need for any event
	if event.Type() == discord.InteractionTypePing {
		client.Logger().Debug("received http interaction ping. responding with pong")
		if err := respondFunc(discord.InteractionResponse{
			Type: discord.InteractionResponseTypePong,
		}); err != nil {
			client.Logger().Error("failed to respond to http interaction ping", slog.Any("err", err))
		}
		return
	}
	handleInteraction(client, -1, -1, respondFunc, event.Interaction)
}
