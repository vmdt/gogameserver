package dtos

import (
	"time"

	"github.com/vmdt/gogameserver/modules/player/application/dtos"
)

type RoomDTO struct {
	ID                string                `json:"id"`
	Status            string                `json:"status"`
	Turn              int                   `json:"turn"`
	WhoWin            int                   `json:"who_win"`
	IsEnded           bool                  `json:"is_ended"`
	BattleshipOptions *BattleshipOptionsDTO `json:"options,omitempty"`
	CreatedAt         *time.Time            `json:"created_at"`
	UpdatedAt         *time.Time            `json:"updated_at"`
}

type PlayerCreateRoomDTO struct {
	Room   RoomDTO         `json:"room"`
	Player *dtos.PlayerDTO `json:"player"`
}

type RoomInformationDTO struct {
	Room    RoomDTO          `json:"room"`
	Players []*RoomPlayerDTO `json:"players"`
}
