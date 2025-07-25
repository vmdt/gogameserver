package commands

import (
	"context"
	"errors"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/player/application/commands"
	player_dtos "github.com/vmdt/gogameserver/modules/player/application/dtos"
	player_room_cmd "github.com/vmdt/gogameserver/modules/room/application/commands/player_room"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/application/events"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type JoinRoomCommand struct {
	Name     string  `json:"name" validate:"required"`
	PlayerId *string `json:"player_id,omitempty"`
	RoomId   string  `json:"room_id" validate:"required"`
	UserId   *string `json:"user_id"`
}

func NewJoinRoomCommand(name string, userId *string, playerId *string, roomId string) *JoinRoomCommand {
	return &JoinRoomCommand{
		Name:     name,
		UserId:   userId,
		PlayerId: playerId,
		RoomId:   roomId,
	}
}

type JoinRoomHandler struct {
	roomRepo  domain.IRoomRepository
	ctx       context.Context
	log       logger.ILogger
	dbContext *infrastructure.RoomDbContext
}

func NewJoinRoomHandler(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository, dbContext *infrastructure.RoomDbContext) *JoinRoomHandler {
	return &JoinRoomHandler{
		roomRepo:  roomRepo,
		ctx:       ctx,
		log:       log,
		dbContext: dbContext,
	}
}

func (h *JoinRoomHandler) Handle(ctx context.Context, command *JoinRoomCommand) (*dtos.RoomPlayerDTO, error) {
	existingRoom, err := h.roomRepo.GetRoomByID(ctx, command.RoomId)
	if err != nil {
		return nil, err
	}
	if existingRoom == nil {
		return nil, errors.New("room not found")
	}

	if existingRoom.Status != "lobby" {
		return nil, errors.New("room is not in lobby status")
	}

	var roomPlayerCount int64
	if err := h.dbContext.GetModelDB(&domain.RoomPlayer{}).Where("room_id = ?", existingRoom.ID).Count(&roomPlayerCount).Error; err != nil {
		return nil, err
	}
	if roomPlayerCount >= 2 {
		return nil, errors.New("room is full")
	}

	player, err := mediatr.Send[*commands.CreatePlayerCommand, *player_dtos.PlayerDTO](ctx, commands.NewCreatePlayerCommand(command.Name, command.UserId))
	if err != nil {
		return nil, err
	}

	roomPlayer, err := mediatr.Send[*player_room_cmd.InternalCreateRoomPlayerCommand, *domain.RoomPlayer](ctx, player_room_cmd.NewInternalCreateRoomPlayerCommand(command.RoomId, player.ID, false, 2))
	if err != nil {
		return nil, err
	}

	joinRoomEvent := events.NewJoinRoomEvent(roomPlayer.RoomId.String(), roomPlayer.PlayerId.String())
	if err := mediatr.Publish[*events.JoinRoomEvent](ctx, joinRoomEvent); err != nil {
		h.log.Error("Failed to publish JoinRoomEvent", "error", err)
		return nil, err
	}

	roomDto := &dtos.RoomDTO{
		ID:        roomPlayer.RoomId.String(),
		Status:    existingRoom.Status,
		CreatedAt: existingRoom.CreatedAt,
		UpdatedAt: existingRoom.UpdatedAt,
	}

	playerDto := &player_dtos.PlayerDTO{
		ID:        player.ID,
		Name:      player.Name,
		UserId:    player.UserId,
		CreatedAt: player.CreatedAt,
		UpdatedAt: player.UpdatedAt,
	}
	return &dtos.RoomPlayerDTO{
		RoomId:         roomPlayer.RoomId.String(),
		PlayerId:       roomPlayer.PlayerId.String(),
		IsReady:        roomPlayer.IsReady,
		IsDisconnected: roomPlayer.IsDisconnected,
		DisconnectedAt: roomPlayer.DisconnectedAt,
		Room:           roomDto,
		Player:         playerDto,
	}, nil
}
