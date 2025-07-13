package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
)

type Room struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Status    string     `gorm:"type:varchar(50)" json:"status"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func NewRoom(status string) *Room {
	now := time.Now()
	return &Room{
		ID:        uuid.New(),
		Status:    status,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func (r *Room) ToDTO() *dtos.RoomDTO {
	return &dtos.RoomDTO{
		ID:        r.ID.String(),
		Status:    r.Status,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
