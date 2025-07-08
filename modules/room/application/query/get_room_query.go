package query

import (
	"context"

	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type GetRoomQuery struct {
	ID string `json:"id"`
}

func NewGetRoomQuery(id string) *GetRoomQuery {
	return &GetRoomQuery{
		ID: id,
	}
}

type GetRoomHandler struct {
	log      logger.ILogger
	ctx      context.Context
	roomRepo domain.IRoomRepository
}

func NewGetRoomHandler(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository) *GetRoomHandler {
	return &GetRoomHandler{
		log:      log,
		ctx:      ctx,
		roomRepo: roomRepo,
	}
}

func (h *GetRoomHandler) Handle(ctx context.Context, query *GetRoomQuery) (*dtos.RoomDTO, error) {
	room, err := h.roomRepo.GetRoomByID(ctx, query.ID)
	if err != nil {
		h.log.Error("Failed to fetch room", "error", err)
		return nil, err
	}
	response := &dtos.RoomDTO{
		ID:        room.ID.String(),
		Status:    room.Status,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
	return response, nil
}
