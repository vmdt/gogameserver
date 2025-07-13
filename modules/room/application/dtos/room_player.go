package dtos

import (
	"time"

	"github.com/vmdt/gogameserver/modules/player/application/dtos"
)

type RoomPlayerDTO struct {
	IsReady        bool       `json:"is_ready"`
	IsDisconnected bool       `json:"is_disconnected"`
	DisconnectedAt *time.Time `json:"disconnected_at,omitempty"`
	IsHost         bool       `json:"is_host"`

	RoomId   string          `json:"room_id"`
	PlayerId string          `json:"player_id"`
	Room     *RoomDTO        `json:"room,omitempty"`
	Player   *dtos.PlayerDTO `json:"player,omitempty"`
}
