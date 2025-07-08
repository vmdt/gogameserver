package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/player/application/dtos"
	"github.com/vmdt/gogameserver/modules/player/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CreatePlayerCommand struct {
	Name   string  `json:"name" validate:"required"`
	UserId *string `json:"user_id"`
}

func NewCreatePlayerCommand(name string, userId *string) *CreatePlayerCommand {
	return &CreatePlayerCommand{
		Name:   name,
		UserId: userId,
	}
}

type CreatePlayerHandler struct {
	log        logger.ILogger
	ctx        context.Context
	playerRepo domain.IPlayerRepository
}

func NewCreatePlayerHandler(log logger.ILogger, ctx context.Context, playerRepo domain.IPlayerRepository) *CreatePlayerHandler {
	return &CreatePlayerHandler{
		log:        log,
		ctx:        ctx,
		playerRepo: playerRepo,
	}
}

func (h *CreatePlayerHandler) Handle(ctx context.Context, command *CreatePlayerCommand) (*dtos.PlayerDTO, error) {
	player, err := h.playerRepo.CreatePlayer(ctx, &domain.Player{
		ID:     uuid.New(),
		Name:   command.Name,
		UserId: command.UserId,
	})
	if err != nil {
		h.log.Error("Failed to create player", "error", err)
		return nil, err
	}
	h.log.Info("Player created successfully", "player_id", player.ID)
	return &dtos.PlayerDTO{
		ID:        player.ID.String(),
		Name:      player.Name,
		UserId:    player.UserId,
		CreatedAt: player.CreatedAt,
		UpdatedAt: player.UpdatedAt,
	}, nil
}
