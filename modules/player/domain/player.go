package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/player/application/dtos"
)

type Player struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	UserId    *string   `gorm:"type:uuid;default:null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (p *Player) ToDTO() *dtos.PlayerDTO {
	return &dtos.PlayerDTO{
		ID:        p.ID.String(),
		Name:      p.Name,
		UserId:    p.UserId,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
