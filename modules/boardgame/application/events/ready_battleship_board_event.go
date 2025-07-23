package events

type ReadyBattleShipBoardEvent struct {
	PlayerId string `json:"player_id"`
	RoomId   string `json:"room_id"`
}

func NewReadyBattleShipBoardEvent(playerId, roomId string) *ReadyBattleShipBoardEvent {
	return &ReadyBattleShipBoardEvent{
		PlayerId: playerId,
		RoomId:   roomId,
	}
}

func (e *ReadyBattleShipBoardEvent) GetPlayerId() string {
	return e.PlayerId
}

func (e *ReadyBattleShipBoardEvent) GetRoomId() string {
	return e.RoomId
}
