package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/boardgame/application/commands"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/modules/boardgame/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigBattleShipMediator(
	log logger.ILogger,
	ctx context.Context,
	redisClient *redis.Client,
	db *infrastructure.BoardGameDbContext,
	battleshipRepo domain.IBattleShipRepository,
) {
	err := mediatr.RegisterRequestHandler(commands.NewCreateBattleShipBoardCommandHandler(log, ctx, battleshipRepo))
	if err != nil {
		log.Fatalf("failed to register command handler: %v", err)
	}
}
