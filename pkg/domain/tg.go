package domain

type TelegramBotToken string

type Update struct {
	UpdateID int
	Message  *Message
}
