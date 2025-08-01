package commands

import (
	"context"
	"errors"

	"github.com/mehdihadeli/go-mediatr"
	boardgame_events "github.com/vmdt/gogameserver/modules/boardgame/application/events"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type SetWhoWinCommand struct {
	RoomId   string `param:"room_id" validate:"required" json:"-"`
	PlayerId string `json:"player_id"`
}

func NewSetWhoWinCommand(roomId, playerId string) *SetWhoWinCommand {
	return &SetWhoWinCommand{
		RoomId:   roomId,
		PlayerId: playerId,
	}
}

type SetWhoWinCommandHandler struct {
	log       logger.ILogger
	ctx       context.Context
	roomRepo  domain.IRoomRepository
	dbContext *infrastructure.RoomDbContext
}

func NewSetWhoWinCommandHandler(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository, dbContext *infrastructure.RoomDbContext) *SetWhoWinCommandHandler {
	return &SetWhoWinCommandHandler{
		log:       log,
		ctx:       ctx,
		roomRepo:  roomRepo,
		dbContext: dbContext,
	}
}

func (h *SetWhoWinCommandHandler) Handle(ctx context.Context, command *SetWhoWinCommand) (*dtos.RoomDTO, error) {
	var roomPlayer *domain.RoomPlayer
	if err := h.dbContext.GetModelDB(&domain.RoomPlayer{}).Where("player_id = ? AND room_id = ?", command.PlayerId, command.RoomId).First(&roomPlayer).Error; err != nil {
		h.log.Error("Failed to get room player", "error", err, "player_id", command.PlayerId, "room_id", command.RoomId)
		return nil, err
	}

	var room *domain.Room
	if err := h.dbContext.GetModelDB(&domain.Room{}).Where("id = ?", command.RoomId).First(&room).Error; err != nil {
		h.log.Error("Failed to get room by ID", "error", err, "room_id", command.RoomId)
		return nil, err
	}

	if room == nil {
		h.log.Error("Room not found", "room_id", command.RoomId)
		return nil, errors.New("room not found")
	}

	room.IsEnded = true
	room.WhoWin = roomPlayer.Me
	if err := h.dbContext.GetModelDB(&domain.Room{}).Where("id = ?", command.RoomId).Updates(room).Error; err != nil {
		h.log.Error("Failed to update room endgame status", "error", err, "room_id", command.RoomId)
		return nil, err
	}
	h.log.Info("Room endgame status updated successfully", "room_id", command.RoomId, "who_win", roomPlayer.Me)

	endgameEvent := boardgame_events.NewEndgameEvent(roomPlayer.Me, command.PlayerId, command.RoomId)
	if err := mediatr.Publish(ctx, endgameEvent); err != nil {
		h.log.Error("Failed to publish endgame event", "error", err)
		return nil, err
	}

	roomDTO := room.ToDTO()

	return roomDTO, nil
}
