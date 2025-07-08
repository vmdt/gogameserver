package domain

import (
	"time"

	"github.com/vmdt/gogameserver/modules/player/domain"
)

type RoomPlayer struct {
	RoomId         string     `gorm:"type:uuid;primaryKey" json:"room_id"`
	PlayerId       string     `gorm:"type:uuid;primaryKey" json:"player_id"`
	IsReady        bool       `gorm:"default:false" json:"is_ready"`
	IsDisconnected bool       `gorm:"default:false" json:"is_disconnected"`
	DisconnectedAt *time.Time `gorm:"default:null" json:"disconnected_at,omitempty"`

	Room   *Room          `gorm:"foreignKey:RoomId;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"room,omitempty"`
	Player *domain.Player `gorm:"foreignKey:PlayerId;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"player,omitempty"`
}
