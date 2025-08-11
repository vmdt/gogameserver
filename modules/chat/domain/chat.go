package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/chat/application/dtos"
)

type GameType int

const (
	Battleship GameType = 1 // 1
)

type Chat struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	GameType GameType  `gorm:"default:1" json:"game_type"` // 1 = Battleship
	RoomId   uuid.UUID `gorm:"type:uuid" json:"room_id"`

	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Messages []ChatMessage `json:"messages,omitempty"`
}

func (c *Chat) ToDTO() *dtos.ChatDTO {
	return &dtos.ChatDTO{
		ID:        c.ID.String(),
		GameType:  int(c.GameType),
		RoomId:    c.RoomId.String(),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Messages:  c.GetMessagesDTO(),
	}
}

func (c *Chat) GetMessages() []ChatMessage {
	return c.Messages
}

func (c *Chat) GetMessagesDTO() []dtos.ChatMessageDTO {
	var messagesDTO []dtos.ChatMessageDTO
	for _, message := range c.Messages {
		messagesDTO = append(messagesDTO, *message.ToDTO())
	}
	return messagesDTO
}

func (c *Chat) AddMessage(message *ChatMessage) {
	c.Messages = append(c.Messages, *message)
}
