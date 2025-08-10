package events

import (
	"context"
	"os"

	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/rabbitmq"
)

type SendChatMessageEvent struct {
	SenderId string `json:"sender_id"`
	RoomId   string `json:"room_id"`
	Content  string `json:"content"`
	IsLog    bool   `json:"is_log"`
}

func NewSendChatMessageEvent(senderId, roomId, content string, isLog bool) *SendChatMessageEvent {
	return &SendChatMessageEvent{
		SenderId: senderId,
		RoomId:   roomId,
		Content:  content,
		IsLog:    isLog,
	}
}

type SendChatMessageEventHandler struct {
	log       logger.ILogger
	ctx       context.Context
	publisher rabbitmq.IPublisher
}

func NewSendChatMessageEventHandler(log logger.ILogger, ctx context.Context, publisher rabbitmq.IPublisher) *SendChatMessageEventHandler {
	return &SendChatMessageEventHandler{
		log:       log,
		ctx:       ctx,
		publisher: publisher,
	}
}

func (h *SendChatMessageEventHandler) Handle(ctx context.Context, event *SendChatMessageEvent) error {
	h.log.Info("SendChatMessageEventHandler: Publishing chat message event", "room_id", event.RoomId, "sender_id", event.SenderId)
	var data = map[string]interface{}{
		"sender_id": event.SenderId,
		"room_id":   event.RoomId,
		"content":   event.Content,
		"is_log":    event.IsLog,
	}

	err := h.publisher.PublishMessage(data, "chat_battleship", os.Getenv("CHAT_QUEUE_KEY"))
	if err != nil {
		h.log.Error("SendChatMessageEventHandler: Failed to publish chat message event", "error", err)
		return err
	}
	h.log.Info("SendChatMessageEventHandler: Chat message event published successfully", "room_id", event.RoomId, "sender_id", event.SenderId)
	return nil
}
