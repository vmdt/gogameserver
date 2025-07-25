package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/player/application/commands"
	player_dtos "github.com/vmdt/gogameserver/modules/player/application/dtos"
	player_room_cmd "github.com/vmdt/gogameserver/modules/room/application/commands/player_room"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/application/events"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type PlayerCreateRoomCommand struct {
	Name   string  `json:"name" validate:"required"`
	UserId *string `json:"user_id"`
}

func NewPlayerCreateRoomCommand(name string, userId *string) *PlayerCreateRoomCommand {
	return &PlayerCreateRoomCommand{
		Name:   name,
		UserId: userId,
	}
}

type PlayerCreateRoomHandler struct {
	roomRepo domain.IRoomRepository
	ctx      context.Context
	log      logger.ILogger
}

func NewPlayerCreateRoomHandler(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository) *PlayerCreateRoomHandler {
	return &PlayerCreateRoomHandler{
		roomRepo: roomRepo,
		ctx:      ctx,
		log:      log,
	}
}

func (h *PlayerCreateRoomHandler) Handle(ctx context.Context, command *PlayerCreateRoomCommand) (*dtos.RoomPlayerDTO, error) {
	player, err := mediatr.Send[*commands.CreatePlayerCommand, *player_dtos.PlayerDTO](ctx, commands.NewCreatePlayerCommand(command.Name, command.UserId))
	if err != nil {
		return nil, err
	}

	room, err := h.roomRepo.CreateRoom(ctx, &domain.Room{
		ID:     uuid.New(),
		Status: "lobby",
	})
	if err != nil {
		return nil, err
	}

	roomPlayer, err := mediatr.Send[*player_room_cmd.InternalCreateRoomPlayerCommand, *domain.RoomPlayer](ctx, player_room_cmd.NewInternalCreateRoomPlayerCommand(room.ID.String(), player.ID, true, 1))
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
		Status:    room.Status,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
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
