package events

type IRoomReadyEvent interface {
	GetRoomId() string
	GetPlayerId() string
}
