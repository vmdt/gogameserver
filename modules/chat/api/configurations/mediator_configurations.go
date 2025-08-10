package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/chat/application/commands"
	"github.com/vmdt/gogameserver/modules/chat/application/queries"
	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/modules/chat/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigChatMediators(
	log logger.ILogger,
	ctx context.Context,
	redisClient *redis.Client,
	db *infrastructure.ChatDbContext,
	chatRepo domain.IChatRepository,
) error {
	// Register command handlers
	err := mediatr.RegisterRequestHandler(
		commands.NewCreateChatCommandHandler(log, ctx, chatRepo),
	)
	if err != nil {
		log.Errorf("failed to register CreateChatCommandHandler: %v", err)
		return err
	}

	// Register queries handlers
	err = mediatr.RegisterRequestHandler(
		queries.NewGetChatQueryHandler(log, ctx, chatRepo),
	)
	if err != nil {
		log.Errorf("failed to register GetChatQueryHandler: %v", err)
		return err
	}
	return nil
}
