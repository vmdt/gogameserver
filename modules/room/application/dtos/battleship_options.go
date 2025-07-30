package dtos

import "time"

type BattleshipOptionsDTO struct {
	Id            string     `json:"id"`
	TimePerTurn   int        `json:"time_per_turn"`   // in seconds
	TimePlaceShip int        `json:"time_place_ship"` // in seconds
	WhoGoFirst    int        `json:"who_go_first"`    // 0: random, 1: player1, 2: player2
	StartPlaceAt  *time.Time `json:"start_place_at"`  // when the player can start placing ships
	RoomId        string     `json:"room_id"`         // UUID of the room
}
