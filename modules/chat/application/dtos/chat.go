package dtos

import "time"

type ChatDTO struct {
	ID        string     `json:"id"`
	GameType  int        `json:"game_type"` // 1 = Battleship
	RoomId    string     `json:"room_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Messages []ChatMessageDTO `json:"messages,omitempty"`
}

type ChatMessageDTO struct {
	ID       string `json:"id"`
	ChatId   string `json:"chat_id"`
	SenderId string `json:"sender_id"`
	Content  string `json:"content"`
	IsLog    bool   `json:"is_log"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
