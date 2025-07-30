package domain

import (
	"time"

	"github.com/google/uuid"
)

type BattleshipOptions struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	TimePerTurn   time.Duration `gorm:"default:30" json:"time_per_turn"`
	TimePlaceShip time.Duration `gorm:"default:120" json:"time_place_ship"`
	WhoGoFirst    int           `gorm:"default:0" json:"who_go_first"` // 0: random, 1: player1, 2: player2
	StartPlaceAt  *time.Time    `gorm:"default:null" json:"start_place_at,omitempty"`

	RoomId uuid.UUID `gorm:"type:uuid;primaryKey" json:"room_id"`
	Room   *Room     `gorm:"foreignKey:RoomId;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"room,omitempty"`
}

func (bo *BattleshipOptions) TableName() string {
	return "room_battleship_options"
}
