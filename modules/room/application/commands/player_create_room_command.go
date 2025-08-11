package commands

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	chat_commands "github.com/vmdt/gogameserver/modules/chat/application/commands"
	chat_dtos "github.com/vmdt/gogameserver/modules/chat/application/dtos"
	"github.com/vmdt/gogameserver/modules/player/application/commands"
	player_dtos "github.com/vmdt/gogameserver/modules/player/application/dtos"
	battleship_options_cmd "github.com/vmdt/gogameserver/modules/room/application/commands/battleship_options"
	player_room_cmd "github.com/vmdt/gogameserver/modules/room/application/commands/player_room"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/application/events"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/server/pkg/auth"
)

type BattleshipOptions struct {
	TimePerTurn   int `json:"time_per_turn"`   // in seconds
	TimePlaceShip int `json:"time_place_ship"` // in seconds
	WhoGoFirst    int `json:"who_go_first"`    // 0: random, 1: player1, 2: player2
}

type PlayerCreateRoomCommand struct {
	Name    string             `json:"name" validate:"required"`
	UserId  *string            `json:"user_id"`
	Options *BattleshipOptions `json:"options"`
}

func NewPlayerCreateRoomCommand(name string, userId *string, options *BattleshipOptions) *PlayerCreateRoomCommand {
	return &PlayerCreateRoomCommand{
		Name:    name,
		UserId:  userId,
		Options: options,
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
	turn := 1
	if command.Options != nil {
		if command.Options.WhoGoFirst == 0 {
			// Randomly select between 1 and 2
			if uuid.New().ID()%2 == 0 {
				turn = 1
			} else {
				turn = 2
			}
		} else {
			turn = command.Options.WhoGoFirst
		}
	}

	room, err := h.roomRepo.CreateRoom(ctx, &domain.Room{
		ID:     uuid.New(),
		Status: "lobby",
		Turn:   turn,
	})
	if err != nil {
		return nil, err
	}

	roomPlayer, err := mediatr.Send[*player_room_cmd.InternalCreateRoomPlayerCommand, *domain.RoomPlayer](ctx, player_room_cmd.NewInternalCreateRoomPlayerCommand(room.ID.String(), player.ID, true, 1))
	if err != nil {
		return nil, err
	}

	var battleshipOptionsDto *dtos.BattleshipOptionsDTO
	if command.Options != nil {
		battleshipOptionsCmd := battleship_options_cmd.NewCreateBattleshipOptionsCmd(
			command.Options.TimePerTurn,
			command.Options.TimePlaceShip,
			command.Options.WhoGoFirst,
			roomPlayer.RoomId.String(),
		)
		var err error
		battleshipOptionsDto, err = mediatr.Send[*battleship_options_cmd.CreateBattleshipOptionsCmd, *dtos.BattleshipOptionsDTO](ctx, battleshipOptionsCmd)
		if err != nil {
			h.log.Error("Failed to create battleship options", "error", err)
			return nil, err
		}
	}

	createRoomChat := chat_commands.NewCreateChatCommand(
		room.ID.String(),
		1, // battleship
	)

	_, err = mediatr.Send[*chat_commands.CreateChatCommand, *chat_dtos.ChatDTO](ctx, createRoomChat)
	if err != nil {
		h.log.Error("Failed to create chat for room", "error", err)
		return nil, err
	}

	userId := auth.GetUserId(ctx)
	if userId == "" {
		h.log.Error("User ID not found in context")
		return nil, errors.New("invalid user ID")
	}
	joinRoomEvent := events.NewJoinRoomEvent(roomPlayer.RoomId.String(), roomPlayer.PlayerId.String(), userId)
	if err := mediatr.Publish[*events.JoinRoomEvent](ctx, joinRoomEvent); err != nil {
		h.log.Error("Failed to publish JoinRoomEvent", "error", err)
		return nil, err
	}

	roomDto := &dtos.RoomDTO{
		ID:                roomPlayer.RoomId.String(),
		Status:            room.Status,
		Turn:              room.Turn,
		BattleshipOptions: battleshipOptionsDto,
		CreatedAt:         room.CreatedAt,
		UpdatedAt:         room.UpdatedAt,
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
		Me:             roomPlayer.Me,
		Room:           roomDto,
		Player:         playerDto,
	}, nil
}
