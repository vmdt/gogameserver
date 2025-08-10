package domain

type IChatRepository interface {
	CreateRoom(chat *Chat) (*Chat, error)
	ChatChatByRoomId(roomId string) (*Chat, error)
}
