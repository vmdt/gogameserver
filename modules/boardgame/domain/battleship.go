package domain

import "github.com/google/uuid"

type BattleShip struct {
	RoomId   uuid.UUID `gorm:"type:uuid;primaryKey" json:"room_id"`
	PlayerId uuid.UUID `gorm:"type:uuid;primaryKey" json:"player_id"`
	Ships    []Ship    `gorm:"type:jsonb" json:"ships"`
	Shots    []Shot    `gorm:"type:jsonb" json:"shots"`
}

func (b *BattleShip) TableName() string {
	return "boardgame_battleship"
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
