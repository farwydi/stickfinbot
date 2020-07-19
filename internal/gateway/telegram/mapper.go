package telegram

import (
	"github.com/farwydi/stickfinbot/pkg/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func fromMessageConfig(m domain.MessageConfig) tgbotapi.Chattable {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           m.BaseChat.ChatID,
			ReplyToMessageID: m.BaseChat.ReplyToMessageID,
		},
		Text: m.Text,
	}
}

func toMessage(m tgbotapi.Message) *domain.Message {
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
