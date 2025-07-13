package commands

import (
	"context"

	"github.com/vmdt/gogameserver/modules/player/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type DeletePlayerCommand struct {
	PlayerId string `json:"player_id"`
}

func NewDeletePlayerCommand(playerId string) *DeletePlayerCommand {
	return &DeletePlayerCommand{
		PlayerId: playerId,
	}
}

type DeletePlayerCommandHandler struct {
	log        logger.ILogger
	ctx        context.Context
	playerRepo domain.IPlayerRepository
}

func NewDeletePlayerCommandHandler(log logger.ILogger, ctx context.Context, playerRepo domain.IPlayerRepository) *DeletePlayerCommandHandler {
	return &DeletePlayerCommandHandler{
		log:        log,
		ctx:        ctx,
		playerRepo: playerRepo,
	}
}

func (h *DeletePlayerCommandHandler) Handle(ctx context.Context, command *DeletePlayerCommand) (bool, error) {
	return h.playerRepo.DeletePlayer(ctx, command.PlayerId)
}
