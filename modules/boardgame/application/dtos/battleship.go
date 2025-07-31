package dtos

import "github.com/vmdt/gogameserver/modules/boardgame/domain"

type BattleshipGame struct {
	PlayerId      string        `json:"player_id"`
	RoomId        string        `json:"room_id"`
	Ships         []domain.Ship `json:"ships"`
	Shots         []domain.Shot `json:"shots"`
	OpponentShots []domain.Shot `json:"opponent_shots"`
}

type WinStatusDTO struct {
	PlayerId string `json:"player_id"`
	Win      bool   `json:"win"`
	Placed   bool   `json:"placed"`
}

type WhoWinDTO struct {
	RoomId    string         `json:"room_id"`
	WinStatus []WinStatusDTO `json:"win_status"`
}
