module github.com/nextep-community/gocord/_examples/application_commands/http

go 1.21

replace github.com/nextep-community/gocord => ../../../

require (
	github.com/nextep-community/gocord v0.18.11
	github.com/disgoorg/snowflake/v2 v2.0.3
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a
)

require (
	github.com/disgoorg/json v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/sasha-s/go-csync v0.0.0-20240107134140-fcbab37b09ad // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
)