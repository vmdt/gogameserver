package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/chat/application/dtos"
	player_domain "github.com/vmdt/gogameserver/modules/player/domain"
)

type ChatMessage struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Content  string    `gorm:"type:text" json:"content"`
	SenderId uuid.UUID `gorm:"type:uuid" json:"sender_id"`
	IsLog    bool      `gorm:"default:false" json:"is_log"`

	Sender player_domain.Player `gorm:"foreignKey:SenderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	ChatId uuid.UUID            `gorm:"type:uuid" json:"chat_id"`
	Chat   Chat                 `gorm:"foreignKey:ChatId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (m *ChatMessage) ToDTO() *dtos.ChatMessageDTO {
	return &dtos.ChatMessageDTO{
		ID:        m.ID.String(),
		ChatId:    m.ChatId.String(),
		SenderId:  m.SenderId.String(),
		Content:   m.Content,
		IsLog:     m.IsLog,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
