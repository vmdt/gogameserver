package dtos

import "time"

type PlayerDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
