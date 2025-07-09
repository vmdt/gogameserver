package events

type JoinRoomEvent struct {
	RoomId   string `json:"room_id"`
	PlayerId string `json:"player_id"`
}

func NewJoinRoomEvent(roomId, playerId string) *JoinRoomEvent {
	return &JoinRoomEvent{
		RoomId:   roomId,
		PlayerId: playerId,
	}
}
