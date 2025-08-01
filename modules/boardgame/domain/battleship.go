package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type BattleShip struct {
	RoomId   uuid.UUID      `gorm:"type:uuid;primaryKey" json:"room_id"`
	PlayerId uuid.UUID      `gorm:"type:uuid;primaryKey" json:"player_id"`
	Ships    datatypes.JSON `gorm:"type:jsonb" json:"ships"`
	Shots    datatypes.JSON `gorm:"type:jsonb" json:"shots"`

	CreatedAt      *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	OpponentShotAt *time.Time `json:"opponent_shot_at,omitempty"`
}

func (b *BattleShip) TableName() string {
	return "boardgame_battleship"
}

func (b *BattleShip) UpdateOpponentShotAt(shotAt *time.Time) {
	if shotAt != nil {
		b.OpponentShotAt = shotAt
	} else {
		b.OpponentShotAt = nil
	}
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Ship struct {
	Name        string     `json:"name"`
	Positions   []Position `json:"positions"`
	Size        int        `json:"size"`
	Orientation string     `json:"orientation"` // "horizontal" or "vertical"
}

type Shot struct {
	Position Position `json:"position"`
	Status   string   `json:"status"` // e.g., "hit", "miss"
}
