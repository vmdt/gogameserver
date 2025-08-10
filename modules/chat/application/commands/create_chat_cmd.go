package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/chat/application/dtos"
	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CreateChatCommand struct {
	RoomId   string `json:"room_id" validate:"required"`
	GameType int    `json:"game_type" default:"1"`
}

func NewCreateChatCommand(roomId string, gameType int) *CreateChatCommand {
	return &CreateChatCommand{
		RoomId:   roomId,
		GameType: gameType,
	}
}

type CreateChatCommandHandler struct {
	log            logger.ILogger
	ctx            context.Context
	chatRepository domain.IChatRepository
}

func NewCreateChatCommandHandler(log logger.ILogger, ctx context.Context, chatRepository domain.IChatRepository) *CreateChatCommandHandler {
	return &CreateChatCommandHandler{
		log:            log,
		ctx:            ctx,
		chatRepository: chatRepository,
	}
}

func (h *CreateChatCommandHandler) Handle(ctx context.Context, command *CreateChatCommand) (*dtos.ChatDTO, error) {
	chat := &domain.Chat{
		ID:       uuid.New(),
		GameType: domain.GameType(command.GameType),
		RoomId:   uuid.MustParse(command.RoomId),
	}

	createdChat, err := h.chatRepository.CreateRoom(chat)
	if err != nil {
		h.log.Error("CreateChatCommandHandler: Failed to create chat room", "error", err)
		return nil, err
	}
	h.log.Info("CreateChatCommandHandler: Chat room created successfully", "chat_id", createdChat.ID, "room_id", command.RoomId)
	return createdChat.ToDTO(), nil
}
