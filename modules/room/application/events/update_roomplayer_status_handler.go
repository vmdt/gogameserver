package events

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateRoomPlayerStatusHandler[T IRoomReadyEvent] struct {
	log         logger.ILogger
	ctx         context.Context
	redisClient *redis.Client
	db          *infrastructure.RoomDbContext
}

func NewUpdateRoomPlayerStatusHandler[T IRoomReadyEvent](log logger.ILogger, ctx context.Context, redisClient *redis.Client, db *infrastructure.RoomDbContext) *UpdateRoomPlayerStatusHandler[T] {
	return &UpdateRoomPlayerStatusHandler[T]{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
		db:          db,
	}
}

func (h *UpdateRoomPlayerStatusHandler[T]) Handle(ctx context.Context, event T) error {
	roomId := event.GetRoomId()
	playerId := event.GetPlayerId()

	dbResult := h.db.GetModelDB(&domain.RoomPlayer{}).
		Where("room_id = ? AND player_id = ?", roomId, playerId).
		Updates(map[string]interface{}{
			"status": domain.ReadyToBattle,
		})
	if dbResult.Error != nil {
		h.log.Error("Failed to update room player status", "error", dbResult.Error)
		return dbResult.Error
	}

	return nil
}
