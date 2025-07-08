package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/player/application/commands"
	"github.com/vmdt/gogameserver/modules/player/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigPlayerMediator(log logger.ILogger, ctx context.Context, playerRepo domain.IPlayerRepository) error {
	err := mediatr.RegisterRequestHandler(commands.NewCreatePlayerHandler(log, ctx, playerRepo))
	if err != nil {
		return err
	}

	log.Info("Player mediator configurations completed successfully")
	return nil
}
