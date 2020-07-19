package domain

type Chat struct {
	ID   int64
	Type string
}

type Message struct {
	MessageID int
	Text      string
	Chat      *Chat
}
