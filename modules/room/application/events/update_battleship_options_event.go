package events

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateBattleshipOptionsEvent struct {
	RoomId string `json:"room_id"`
}

type UpdateBattleshipOptionsEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	db          *infrastructure.RoomDbContext
	redisClient *redis.Client
}

func NewUpdateBattleshipOptionsEventHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext, redisClient *redis.Client) *UpdateBattleshipOptionsEventHandler {
	return &UpdateBattleshipOptionsEventHandler{
		log:         log,
		ctx:         ctx,
		db:          db,
		redisClient: redisClient,
	}
}

func (h *UpdateBattleshipOptionsEventHandler) Handle(ctx context.Context, event *UpdateBattleshipOptionsEvent) error {
	var battleshipOptions domain.BattleshipOptions
	if err := h.db.GetModelDB(&domain.BattleshipOptions{}).Where("room_id = ?", event.RoomId).First(&battleshipOptions).Error; err != nil {
		h.log.Error("Failed to fetch battleship options", "error", err, "room_id", event.RoomId)
		return err
	}

	// push the event to Redis
	jsonOpts, err := json.Marshal(battleshipOptions)
	if err != nil {
		h.log.Error("Failed to marshal battleship options", "error", err, "room_id", event.RoomId)
		return err
	}
	redisEvent := map[string]string{
		"options": string(jsonOpts),
		"room_id": event.RoomId,
		"event":   "room:update_options",
	}
	data, err := json.Marshal(redisEvent)
	if err != nil {
		h.log.Error("Failed to marshal Redis event", "error", err, "room_id", event.RoomId)
		return err
	}
	if err := h.redisClient.Publish(h.ctx, "room_events", data).Err(); err != nil {
		h.log.Error("Failed to publish Redis event", "error", err, "room_id", event.RoomId)
		return err
	}

	return nil
}
