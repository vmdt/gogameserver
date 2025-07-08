package events

type CreateRoomEvent struct {
	RoomId   string `json:"room_id"`
	PlayerId string `json:"player_id"`
}

func NewCreateRoomEvent(roomId, playerId string) *CreateRoomEvent {
	return &CreateRoomEvent{
		RoomId:   roomId,
		PlayerId: playerId,
	}
}
