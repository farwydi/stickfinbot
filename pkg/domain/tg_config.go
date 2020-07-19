package domain

type BaseChat struct {
	ChatID           int64
	ReplyToMessageID int
}

type MessageConfig struct {
	BaseChat
	Text string
}
