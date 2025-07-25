package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	boardgame_events "github.com/vmdt/gogameserver/modules/boardgame/application/events"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/utils"
)

type AttackBattleShipBoardEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	db          *infrastructure.RoomDbContext
	redisClient *redis.Client
}

func NewAttackBattleShipBoardEventHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext, redisClient *redis.Client) *AttackBattleShipBoardEventHandler {
	return &AttackBattleShipBoardEventHandler{
		log:         log,
		ctx:         ctx,
		db:          db,
		redisClient: redisClient,
	}
}

func (h *AttackBattleShipBoardEventHandler) Handle(ctx context.Context, event *boardgame_events.AttackBattleShipBoardEvent) error {
	go func(db *infrastructure.RoomDbContext, event *boardgame_events.AttackBattleShipBoardEvent) {
		h.log.Info("Handling AttackBattleShipBoardEvent", "player_id", event.PlayerId, "room_id", event.RoomId, "shot", event.Shot)
		var roomPlayers []domain.RoomPlayer
		if err := db.GetModelDB(&domain.RoomPlayer{}).Where("room_id = ?", event.RoomId).Find(&roomPlayers).Error; err != nil {
			h.log.Error("Failed to get room players", "error", err)
			return
		}

		opponentPlayer := utils.Filter(roomPlayers, func(player domain.RoomPlayer) bool {
			return player.PlayerId != uuid.MustParse(event.PlayerId)
		})

		if len(opponentPlayer) == 0 {
			h.log.Error("Opponent player not found", "room_id", event.RoomId, "player_id", event.PlayerId)
			return
		}

		// update turn in the room
		if err := db.GetModelDB(&domain.Room{}).Where("id = ?", event.RoomId).Update("turn", opponentPlayer[0].Me).Error; err != nil {
			h.log.Error("Failed to update room turn", "error", err)
			return
		}
	}(h.db, event)
	return nil
}
