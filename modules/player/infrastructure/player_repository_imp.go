package infrastructure

import (
	"context"

	"github.com/vmdt/gogameserver/modules/player/domain"
)

type PlayerRepositoryImp struct {
	playerDbContext *PlayerDbContext
}

func NewPlayerRepository(playerDbContext *PlayerDbContext) domain.IPlayerRepository {
	return &PlayerRepositoryImp{
		playerDbContext: playerDbContext,
	}
}

func (p *PlayerRepositoryImp) CreatePlayer(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	if err := p.playerDbContext.GetModelDB().Create(player).Error; err != nil {
		return nil, err
	}
	return player, nil
}

func (p *PlayerRepositoryImp) GetPlayerByID(ctx context.Context, id string) (*domain.Player, error) {
	var player domain.Player
	if err := p.playerDbContext.GetModelDB().First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
