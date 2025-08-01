package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
)

type Room struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Status    string     `gorm:"type:varchar(50)" json:"status"`
	Turn      int        `gorm:"default:0" json:"turn"`
	WhoWin    int        `gorm:"default:0" json:"who_win"`
	IsEnded   bool       `gorm:"default:false" json:"is_ended"`
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
		Turn:      r.Turn,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
