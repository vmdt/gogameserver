package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/chat/application/commands"
	"github.com/vmdt/gogameserver/modules/chat/application/events"
	"github.com/vmdt/gogameserver/modules/chat/application/queries"
	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/modules/chat/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/rabbitmq"
)

func ConfigChatMediators(
	log logger.ILogger,
	ctx context.Context,
	redisClient *redis.Client,
	publisher rabbitmq.IPublisher,
	db *infrastructure.ChatDbContext,
	chatRepo domain.IChatRepository,
	chatMsgRepo domain.IChatMessageRepository,
) error {
	// Register command handlers
	err := mediatr.RegisterRequestHandler(
		commands.NewCreateChatCommandHandler(log, ctx, chatRepo),
	)
	if err != nil {
		log.Errorf("failed to register CreateChatCommandHandler: %v", err)
		return err
	}

	err = mediatr.RegisterRequestHandler(
		commands.NewChatMessageCommandHandler(log, ctx, chatRepo, chatMsgRepo),
	)
	if err != nil {
		log.Errorf("failed to register ChatMessageCommandHandler: %v", err)
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

	// Register event handlers
	err = mediatr.RegisterNotificationHandler(
		events.NewSendChatMessageEventHandler(log, ctx, publisher),
	)
	if err != nil {
		log.Errorf("failed to register SendChatMessageEventHandler: %v", err)
		return err
	}

	return nil
}
