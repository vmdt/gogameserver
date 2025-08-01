package dtos

import (
	"time"

	"github.com/vmdt/gogameserver/modules/boardgame/domain"
)

type BattleshipGame struct {
	PlayerId       string        `json:"player_id"`
	RoomId         string        `json:"room_id"`
	Ships          []domain.Ship `json:"ships"`
	Shots          []domain.Shot `json:"shots"`
	OpponentShots  []domain.Shot `json:"opponent_shots"`
	CreatedAt      *time.Time    `json:"created_at"`
	UpdatedAt      *time.Time    `json:"updated_at"`
	OpponentShotAt *time.Time    `json:"opponent_shot_at,omitempty"`
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

type SunkShipDTO struct {
	ShipName string `json:"ship_name"`
	Size     int    `json:"size"`
	IsSunk   bool   `json:"is_sunk"`
}

type SunkShipsDTO struct {
	PlayerId string        `json:"player_id"`
	Ships    []SunkShipDTO `json:"ships"`
}
