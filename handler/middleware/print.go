package middleware

import (
	"github.com/nextep-community/gocord/handler"
)

func Print(content string) handler.Middleware {
	return func(next handler.Handler) handler.Handler {
		return func(event *handler.InteractionEvent) error {
			println(content)
			return next(event)
		}
	}
}
