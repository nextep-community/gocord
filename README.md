[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/nextep-community/gocord/blob/master/LICENSE).
[![Go Reference](https://pkg.go.dev/badge/github.com/nextep-community/gocord.svg)](https://pkg.go.dev/github.com/nextep-community/gocord)
[![Go Report](https://goreportcard.com/badge/github.com/nextep-community/gocord)](https://goreportcard.com/report/github.com/nextep-community/gocord)
[![Go Version](https://img.shields.io/github/go-mod/go-version/disgoorg/disgo)](https://golang.org/doc/devel/release.html)

# Gocord

Gocord is a Go library forked from [DisGo](https://github.com/disgoorg/disgo),
that is a Discord API wrapper written in Go.

> **Why fork?:** This project was created for internal use within the bot
> [Digo](https://github.com/nextep-community/digo), which was developed by
> Nextep to study and enhance the team's knowledge of Go. For this reason, we
> wanted maximum control and to ensure the library is updated as quickly as
> possible with the latest Discord changes. However, we didn’t want to start
> from scratch.

### Stability

The public API of Gocord is mostly stable at this point in time. Smaller
breaking changes can happen before the v1 is released.

After v1 is released breaking changes may only happen if the Discord API
requires them. They tend to break their released API versions now and then. In
general for every new Discord API version the major version of Gocord should be
increased and with that breaking changes between non-major versions should be
held to a minimum.

## Documentation

Documentation is wip and can be found under

- [![Go Reference](https://pkg.go.dev/badge/github.com/nextep-community/gocord.svg)](https://pkg.go.dev/github.com/nextep-community/gocord)
- [![Discord Documentation](https://img.shields.io/badge/Discord%20Documentation-blue.svg)](https://discord.com/developers/docs)
- [Examples](./_examples/README.md)

GitHub Wiki is currently under construction. We appreciate help here.

### Features

- Full Rest API coverage
- [Gateway](https://discord.com/developers/docs/topics/gateway)
- [Sharding](https://discord.com/developers/docs/topics/gateway#sharding)
- [HTTP Interactions](https://discord.com/developers/docs/interactions/slash-commands#receiving-an-interaction)
- [Application Commands](https://discord.com/developers/docs/interactions/application-commands)
- [Message Components](https://discord.com/developers/docs/interactions/message-components)
- [Modals](https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-modal)
- [Stage Instance](https://discord.com/developers/docs/resources/stage-instance)
- [Guild Template](https://discord.com/developers/docs/resources/guild-template)
- [Sticker](https://discord.com/developers/docs/resources/sticker)
- [RateLimit](https://discord.com/developers/docs/topics/rate-limits)
- [Webhook](https://discord.com/developers/docs/resources/webhook)
- [OAuth2](https://discord.com/developers/docs/topics/oauth2)
- [Threads](https://discord.com/developers/docs/topics/threads)
- [Guild Scheduled Event](https://discord.com/developers/docs/resources/guild-scheduled-event)
- [Voice](https://discord.com/developers/docs/topics/voice-connections)

### Missing Features

- [RPC](https://discord.com/developers/docs/topics/rpc)
  (https://github.com/nextep-community/gocord/pull/170)

## Getting Started

### Installing

```sh
$ go get github.com/nextep-community/gocord
```

### Building a gocord Instance

Build a bot client to interact with the Discord API

```go
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nextep-community/gocord"
	"github.com/nextep-community/gocord/bot"
	"github.com/nextep-community/gocord/events"
	"github.com/nextep-community/gocord/gateway"
)

func main() {
	client, err := gocord.New("token",
		// set gateway options
		bot.WithGatewayConfigOpts(
			// set enabled intents
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
		// add event listeners
		bot.WithEventListenerFunc(func(e *events.MessageCreate) {
			// event code here
		}),
	)
	if err != nil {
		panic(err)
	}
	// connect to the gateway
	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
```

### Logging

gocord uses [slog](https://pkg.go.dev/log/slog) for logging.

## Examples

You can find examples
[here](https://github.com/nextep-community/gocord/tree/master/_examples)

There is also a bot template with commands & db
[here](https://github.com/disgoorg/bot-template)

## Other interesting projects

### [Lavalink](https://github.com/freyacodes/Lavalink)

Is a standalone audio sending node based on
[Lavaplayer](https://github.com/sedmelluq/lavaplayer) and JDA-Audio. Which
allows for sending audio without it ever reaching any of your shards. Lavalink
can be used in combination with

### [Golink](https://github.com/nextep-community/golink)

Is a [Lavalink-Client](https://github.com/freyacodes/Lavalink) which can be used
to communicate with Lavalink to play/search tracks

## Other Golang Discord Libraries

- [DisGo](https://github.com/disgoorg/disgo)
- [discordgo](https://github.com/bwmarrin/discordgo)
- [disgord](https://github.com/andersfylling/disgord)
- [arikawa](https://github.com/diamondburned/arikawa)
- [corde](https://github.com/Karitham/corde)

## Troubleshooting

TODO

## Contributing

TODO

## Special Thanks

We would like to thank the [DisGo](https://github.com/disgoorg/disgo) team and
community.

## License

See LICENSE for more information.

[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/nextep-community/gocord/blob/master/LICENSE).
