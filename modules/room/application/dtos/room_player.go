package dtos

import (
	"time"

	"github.com/vmdt/gogameserver/modules/player/application/dtos"
)

// type RoomPlayerStatus int

// const (
// 	InLobby       RoomPlayerStatus = iota // 0
// 	Placing                               // 1
// 	ReadyToBattle                         // 2
// )

type RoomPlayerDTO struct {
	IsReady        bool       `json:"is_ready"`
	IsDisconnected bool       `json:"is_disconnected"`
	DisconnectedAt *time.Time `json:"disconnected_at,omitempty"`
	IsHost         bool       `json:"is_host"`
	Status         int        `json:"status"`
	Me             int        `json:"me"`

	RoomId   string          `json:"room_id"`
	PlayerId string          `json:"player_id"`
	Room     *RoomDTO        `json:"room,omitempty"`
	Player   *dtos.PlayerDTO `json:"player,omitempty"`
}
