package general

import (
	"context"
	"errors"
	"github.com/farwydi/stickfinbot/pkg/domain"
)

func NewGeneralUseCase(tg domain.TelegramGateway, us domain.UserStatGateway) *General {
	return &General{
		tg: tg,
		us: us,
	}
}

type General struct {
	tg domain.TelegramGateway
	us domain.UserStatGateway
}

func (g *General) Proc(ctx context.Context, update *domain.Update) error {
	// ignore any non-Message Updates
	if update.Message == nil {
		return errors.New("not support method")
	}

	msg := domain.MessageConfig{
		BaseChat: domain.BaseChat{
			ChatID:           update.Message.Chat.ID,
			ReplyToMessageID: update.Message.MessageID,
		},
		Text: update.Message.Text,
	}
	_, err := g.tg.SendMessage(ctx, msg)
	return err
}
