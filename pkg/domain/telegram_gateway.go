package domain

import "context"

type TelegramGateway interface {
	SendMessage(context.Context, MessageConfig) (*Message, error)
}
