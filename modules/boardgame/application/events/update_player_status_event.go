package events

type UpdatePlayerStatusEvent struct {
	PlayerId string `json:"player_id"`
	RoomId   string `json:"room_id"`
}

func NewUpdatePlayerStatusEvent(playerId, roomId string) *UpdatePlayerStatusEvent {
	return &UpdatePlayerStatusEvent{
		PlayerId: playerId,
		RoomId:   roomId,
	}
}

func (e *UpdatePlayerStatusEvent) GetPlayerId() string {
	return e.PlayerId
}

func (e *UpdatePlayerStatusEvent) GetRoomId() string {
	return e.RoomId
}
