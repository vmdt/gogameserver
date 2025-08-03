package events

type JoinRoomEvent struct {
	RoomId   string `json:"room_id"`
	PlayerId string `json:"player_id"`
	UserId   string `json:"user_id"` // Added user_id to track the user joining the room
}

func NewJoinRoomEvent(roomId, playerId, userId string) *JoinRoomEvent {
	return &JoinRoomEvent{
		RoomId:   roomId,
		PlayerId: playerId,
		UserId:   userId,
	}
}
