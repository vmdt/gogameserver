package dtos

import (
	"encoding/json"
	"time"
)

type UserDTO struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Nation    string     `json:"nation"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type TokenPairDTO struct {
	AccessToken  json.RawMessage `json:"access_token"`
	RefreshToken json.RawMessage `json:"refresh_token"`
	ExpiresIn    time.Duration   `json:"expires_in" swaggertype:"integer"`
}

type UserAuthDTO struct {
	User      *UserDTO      `json:"user"`
	TokenPair *TokenPairDTO `json:"tokens"`
}
