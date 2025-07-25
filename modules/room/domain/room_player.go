package domain

import (
	"time"

	"github.com/google/uuid"
	player_dtos "github.com/vmdt/gogameserver/modules/player/application/dtos"
	"github.com/vmdt/gogameserver/modules/player/domain"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
)

type RoomPlayer struct {
	RoomId         uuid.UUID        `gorm:"type:uuid;primaryKey" json:"room_id"`
	PlayerId       uuid.UUID        `gorm:"type:uuid;primaryKey" json:"player_id"`
	IsReady        bool             `gorm:"default:false" json:"is_ready"`
	IsDisconnected bool             `gorm:"default:false" json:"is_disconnected"`
	DisconnectedAt *time.Time       `gorm:"default:null" json:"disconnected_at,omitempty"`
	IsHost         bool             `gorm:"default:false" json:"is_host"`
	Status         RoomPlayerStatus `gorm:"default:0" json:"status"`
	Me             int              `gorm:"default:0" json:"me"`

	Room   *Room          `gorm:"foreignKey:RoomId;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"room,omitempty"`
	Player *domain.Player `gorm:"foreignKey:PlayerId;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"player,omitempty"`
}

type RoomPlayerStatus int

const (
	InLobby       RoomPlayerStatus = iota // 0
	Placing                               // 1
	ReadyToBattle                         // 2
)

func (rp *RoomPlayer) ToDTO() *dtos.RoomPlayerDTO {
	var roomDTO *dtos.RoomDTO
	var playerDTO *player_dtos.PlayerDTO

	if rp.Room != nil {
		roomDTO = rp.Room.ToDTO()
	}
	if rp.Player != nil {
		playerDTO = rp.Player.ToDTO()
	}

	return &dtos.RoomPlayerDTO{
		IsReady:        rp.IsReady,
		IsDisconnected: rp.IsDisconnected,
		DisconnectedAt: rp.DisconnectedAt,
		IsHost:         rp.IsHost,
		Status:         int(rp.Status),
		Me:             rp.Me,
		RoomId:         rp.RoomId.String(),
		PlayerId:       rp.PlayerId.String(),
		Room:           roomDTO,
		Player:         playerDTO,
	}
}
