package commands

import (
	"context"

	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CreateRoomCommand struct {
	Status string `json:"status"`
}

func NewCreateRoomCommand(status string) *CreateRoomCommand {
	return &CreateRoomCommand{
		Status: status,
	}
}

type CreateRoomHandler struct {
	log      logger.ILogger
	ctx      context.Context
	roomRepo domain.IRoomRepository
}

func NewCreateRoomHandler(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository) *CreateRoomHandler {
	return &CreateRoomHandler{
		log:      log,
		ctx:      ctx,
		roomRepo: roomRepo,
	}
}

func (h *CreateRoomHandler) Handle(ctx context.Context, command *CreateRoomCommand) (*dtos.RoomDTO, error) {
	room := domain.NewRoom(command.Status)
	r, err := h.roomRepo.CreateRoom(ctx, room)
	if err != nil {
		h.log.Error("Failed to create room", "error", err)
		return nil, err
	}
	h.log.Info("Room created successfully", "room_id", r.ID)
	response := &dtos.RoomDTO{
		ID:        r.ID.String(),
		Status:    r.Status,
		CreatedAt: r.CreatedAt.String(),
		UpdatedAt: r.UpdatedAt.String(),
	}
	return response, nil
}
