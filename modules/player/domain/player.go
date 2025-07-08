package domain

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	UserId    *string   `gorm:"type:uuid;default:null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
