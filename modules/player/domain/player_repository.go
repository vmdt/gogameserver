package domain

import "context"

type IPlayerRepository interface {
	CreatePlayer(ctx context.Context, player *Player) (*Player, error)
	GetPlayerByID(ctx context.Context, id string) (*Player, error)
}
