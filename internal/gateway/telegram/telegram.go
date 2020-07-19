package telegram

import (
	"context"
	"github.com/farwydi/stickfinbot/pkg/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func NewTelegramGateway(token domain.TelegramBotToken) (domain.TelegramGateway, error) {
	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		return nil, err
	}

	return &tgGateway{
		bot: bot,
	}, nil
}

type tgGateway struct {
	bot *tgbotapi.BotAPI
}

func (t *tgGateway) SendMessage(_ context.Context, msg domain.MessageConfig) (*domain.Message, error) {
	returnedMessage, err := t.bot.Send(fromMessageConfig(msg))
	if err != nil {
		return nil, err
	}

	return toMessage(returnedMessage), nil
}
