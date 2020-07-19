package telegram

import (
	"github.com/farwydi/stickfinbot/pkg/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func toUpdate(m tgbotapi.Update) *domain.Update {
	return &domain.Update{
		UpdateID: m.UpdateID,
		Message:  toMessage(m.Message),
	}
}

func toMessage(m *tgbotapi.Message) *domain.Message {
	return &domain.Message{
		MessageID: m.MessageID,
		Chat:      toChat(m.Chat),
		Text:      m.Text,
	}
}

func toChat(m *tgbotapi.Chat) *domain.Chat {
	return &domain.Chat{
		ID:   m.ID,
		Type: m.Type,
	}
}
