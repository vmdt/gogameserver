package infrastructure

import (
	"context"

	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
)

type ChatRepositoryImp struct {
	chatDbContext *ChatDbContext
	generic       *postgresgorm.GenericRepository[domain.Chat]
	ctx           context.Context
}

func NewChatRepositoryImp(chatDbContext *ChatDbContext) domain.IChatRepository {
	repo := postgresgorm.NewGenericRepository[domain.Chat](chatDbContext.GetModelDB(&domain.Chat{}))
	return &ChatRepositoryImp{
		chatDbContext: chatDbContext,
		generic:       repo,
		ctx:           chatDbContext.context,
	}
}

func (r *ChatRepositoryImp) CreateRoom(chat *domain.Chat) (*domain.Chat, error) {
	if err := r.generic.Add(chat, r.ctx); err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *ChatRepositoryImp) ChatChatByRoomId(roomId string) (*domain.Chat, error) {
	var chat domain.Chat
	if err := r.chatDbContext.GetModelDB(&chat).
		Preload("Messages").
		Where("room_id = ?", roomId).
		First(&chat).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}
