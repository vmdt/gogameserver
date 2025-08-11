package domain

type IChatMessageRepository interface {
	CreateChatMessage(message *ChatMessage) (*ChatMessage, error)
}
