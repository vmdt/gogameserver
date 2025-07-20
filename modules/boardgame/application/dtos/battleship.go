package dtos

import "github.com/vmdt/gogameserver/modules/boardgame/domain"

type BattleshipGame struct {
	PlayerId string        `json:"player_id"`
	RoomId   string        `json:"room_id"`
	Ships    []domain.Ship `json:"ships"`
	Shots    []domain.Shot `json:"shots"`
}
