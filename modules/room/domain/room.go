package domain

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Status    string    `gorm:"type:varchar(50)" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func NewRoom(status string) *Room {
	return &Room{
		ID:        uuid.New(),
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
