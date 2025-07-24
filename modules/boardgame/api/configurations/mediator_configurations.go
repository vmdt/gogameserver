package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/boardgame/application/commands"
	"github.com/vmdt/gogameserver/modules/boardgame/application/events"
	"github.com/vmdt/gogameserver/modules/boardgame/application/queries"
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
	// Register commands mediators
	err := mediatr.RegisterRequestHandler(commands.NewCreateBattleShipBoardCommandHandler(log, ctx, battleshipRepo))
	if err != nil {
		log.Fatalf("failed to register command handler: %v", err)
	}

	err = mediatr.RegisterRequestHandler(commands.NewAttackBattleShipCommandHandler(log, ctx, battleshipRepo))
	if err != nil {
		log.Fatalf("failed to register command handler: %v", err)
	}

	// Register queries mediators
	err = mediatr.RegisterRequestHandler(queries.NewGetBattleshipBoardQueryHandler(log, ctx, battleshipRepo))
	if err != nil {
		log.Fatalf("failed to register query handler: %v", err)
	}

	// Register events mediators
	err = mediatr.RegisterNotificationHandler(events.NewAttackBattleShipBoardEventHandler(log, ctx, redisClient))
	if err != nil {
		log.Fatalf("failed to register event handler: %v", err)
	}
}
