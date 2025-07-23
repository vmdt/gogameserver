package commands

import (
	"context"
	"errors"

	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateRoomStatusCommand struct {
	RoomId string `json:"room_id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

func NewUpdateRoomStatusCommand(roomId, status string) *UpdateRoomStatusCommand {
	return &UpdateRoomStatusCommand{
		RoomId: roomId,
		Status: status,
	}
}

type UpdateRoomStatusCommandHandler struct {
	log      logger.ILogger
	ctx      context.Context
	roomRepo domain.IRoomRepository
}

func NewUpdateRoomStatusCommandHandler(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository) *UpdateRoomStatusCommandHandler {
	return &UpdateRoomStatusCommandHandler{
		log:      log,
		ctx:      ctx,
		roomRepo: roomRepo,
	}
}

func (h *UpdateRoomStatusCommandHandler) Handle(ctx context.Context, command *UpdateRoomStatusCommand) (*dtos.RoomDTO, error) {
	room, err := h.roomRepo.GetRoomByID(ctx, command.RoomId)
	if err != nil {
		h.log.Error("Failed to get room by ID", "error", err)
		return nil, err
	}
	if room == nil {
		h.log.Error("Room not found", "room_id", command.RoomId)
		return nil, errors.New("room not found")
	}
	room.Status = command.Status
	h.log.Info("Updating room status", "room_id", room.ID, "status", command.Status)

	r, err := h.roomRepo.UpdateRoom(ctx, room)
	if err != nil {
		h.log.Error("Failed to update room", "error", err)
		return nil, err
	}
	h.log.Info("Room updated successfully", "room_id", r.ID)
	response := &dtos.RoomDTO{
		ID:        r.ID.String(),
		Status:    r.Status,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
	return response, nil
}
