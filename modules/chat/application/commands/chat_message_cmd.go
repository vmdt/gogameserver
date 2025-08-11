package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/chat/application/dtos"
	"github.com/vmdt/gogameserver/modules/chat/application/events"
	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type ChatMessageCommand struct {
	SenderId string `json:"sender_id" validate:"required"`
	RoomId   string `param:"room_id" validate:"required" json:"-"`
	Content  string `json:"content" validate:"required"`
	IsLog    bool   `json:"is_log" default:"false"`
}

func NewChatMessageCommand(senderId, roomId, content string, isLog bool) *ChatMessageCommand {
	return &ChatMessageCommand{
		SenderId: senderId,
		RoomId:   roomId,
		Content:  content,
		IsLog:    isLog,
	}
}

type ChatMessageCommandHandler struct {
	log                   logger.ILogger
	ctx                   context.Context
	chatRepository        domain.IChatRepository
	chatMessageRepository domain.IChatMessageRepository
}

func NewChatMessageCommandHandler(log logger.ILogger, ctx context.Context, chatRepository domain.IChatRepository, chatMessageRepository domain.IChatMessageRepository) *ChatMessageCommandHandler {
	return &ChatMessageCommandHandler{
		log:                   log,
		ctx:                   ctx,
		chatRepository:        chatRepository,
		chatMessageRepository: chatMessageRepository,
	}
}

func (h *ChatMessageCommandHandler) Handle(ctx context.Context, command *ChatMessageCommand) (*dtos.ChatDTO, error) {
	chat, err := h.chatRepository.ChatChatByRoomId(command.RoomId, true)
	if err != nil {
		h.log.Error("ChatMessageCommandHandler: Failed to get chat by room ID", "room_id", command.RoomId, "error", err)
		return nil, err
	}

	if chat == nil {
		h.log.Error("ChatMessageCommandHandler: No chat found for the provided room ID", "room_id", command.RoomId)
		return nil, nil
	}

	newMessage := &domain.ChatMessage{
		ID:       uuid.New(),
		Content:  command.Content,
		SenderId: uuid.MustParse(command.SenderId),
		IsLog:    command.IsLog,
		ChatId:   chat.ID,
	}

	createdMessage, err := h.chatMessageRepository.CreateChatMessage(newMessage)
	if err != nil {
		h.log.Error("ChatMessageCommandHandler: Failed to create chat message", "error", err)
		return nil, err
	}

	chat.AddMessage(createdMessage)

	// push events - todo
	sendMessageEvent := events.NewSendChatMessageEvent(
		command.SenderId,
		command.RoomId,
		command.Content,
		command.IsLog,
	)
	if err := mediatr.Publish(h.ctx, sendMessageEvent); err != nil {
		h.log.Error("ChatMessageCommandHandler: Failed to handle SendChatMessageEvent", "error", err)
		return nil, err
	}

	return chat.ToDTO(), nil
}
