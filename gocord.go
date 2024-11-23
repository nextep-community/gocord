// Package gocord is a collection of packages for interaction with the Discord Bot and OAuth2 API.
//
// # Discord
//
// Package discord is a collection of structs and types of the Discord API.
//
// # Bot
//
// Package bot connects the Gateway/Sharding, HTTPServer, Cache, Rest & Events packages into a single high level client interface.
//
// # Gateway
//
// Package gateway is used to connect and interact with the Discord Gateway.
//
// # Sharding
//
// Package sharding is used to connect and interact with the Discord Gateway.
//
// # Cache
//
// Package cache provides a generic cache interface for Discord entities.
//
// # HTTPServer
//
// Package httpserver is used to interact with the Discord outgoing webhooks for interactions.
//
// # Events
//
// Package events provide high level events around the Discord Events.
//
// # Rest
//
// Package rest is used to interact with the Discord REST API.
//
// # Webhook
//
// Package webhook provides a high level client interface for interacting with Discord webhooks.
//
// # OAuth2
//
// Package oauth2 provides a high level client interface for interacting with Discord oauth2.
//
// # Voice
//
// Package voice provides a high level client interface for interacting with Discord voice.
package gocord

import (
	"runtime"
	"runtime/debug"

	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/handlers"
)

const (
	// Name is the library name
	Name = "gocord"
	// Module is the library module name
	Module = "github.com/nextep-community/gocord"
	// GitHub is a link to the libraries GitHub repository
	GitHub = "https://github.com/nextep-community/gocord"
)

var (
	// Version is the currently used version of gocord
	Version = getVersion()

	SemVersion = "semver:" + Version
)

func getVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range bi.Deps {
			if dep.Path == Module {
				return dep.Version
			}
		}
	}
	return "unknown"
}

// New creates a new bot.Client with the provided token & bot.ConfigOpt(s)
func New(token string, opts ...bot.ConfigOpt) (bot.Client, error) {
	config := bot.DefaultConfig(handlers.GetGatewayHandlers(), handlers.GetHTTPServerHandler())
	config.Apply(opts)

	return bot.BuildClient(token,
		config,
		handlers.DefaultGatewayEventHandlerFunc,
		handlers.DefaultHTTPServerEventHandlerFunc,
		runtime.GOOS,
		Name,
		GitHub,
		Version,
	)
}
