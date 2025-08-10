package queries

import (
	"context"

	"github.com/vmdt/gogameserver/modules/chat/application/dtos"
	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type GetChatQuery struct {
	RoomId string `param:"room_id" validate:"required" json:"-"`
}

func NewGetChatQuery(roomId string) *GetChatQuery {
	return &GetChatQuery{
		RoomId: roomId,
	}
}

type GetChatQueryHandler struct {
	log            logger.ILogger
	ctx            context.Context
	chatRepository domain.IChatRepository
}

func NewGetChatQueryHandler(log logger.ILogger, ctx context.Context, chatRepository domain.IChatRepository) *GetChatQueryHandler {
	return &GetChatQueryHandler{
		log:            log,
		ctx:            ctx,
		chatRepository: chatRepository,
	}
}

func (h *GetChatQueryHandler) Handle(ctx context.Context, query *GetChatQuery) (*dtos.ChatDTO, error) {
	h.log.Info("GetChatQueryHandler: Fetching chat by room ID", "room_id", query.RoomId)

	chat, err := h.chatRepository.ChatChatByRoomId(query.RoomId)
	if err != nil {
		h.log.Error("GetChatQueryHandler: Failed to fetch chat by room ID", "error", err)
		return nil, err
	}

	if chat == nil {
		h.log.Warn("GetChatQueryHandler: No chat found for the provided room ID", "room_id", query.RoomId)
		return nil, nil
	}

	h.log.Info("GetChatQueryHandler: Chat fetched successfully", "chat_id", chat.ID, "room_id", query.RoomId)
	return chat.ToDTO(), nil
}
