package infrastructure

import (
	"context"

	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
)

type ChatMessageRepositoryImp struct {
	chatMessageDbContext *ChatDbContext
	generic              *postgresgorm.GenericRepository[domain.ChatMessage]
	ctx                  context.Context
}

func NewChatMessageRepositoryImp(chatMessageDbContext *ChatDbContext) domain.IChatMessageRepository {
	repo := postgresgorm.NewGenericRepository[domain.ChatMessage](chatMessageDbContext.GetModelDB(&domain.ChatMessage{}))
	return &ChatMessageRepositoryImp{
		chatMessageDbContext: chatMessageDbContext,
		generic:              repo,
		ctx:                  chatMessageDbContext.context,
	}
}

func (r *ChatMessageRepositoryImp) CreateChatMessage(message *domain.ChatMessage) (*domain.ChatMessage, error) {
	if err := r.generic.Add(message, r.ctx); err != nil {
		return nil, err
	}
	return message, nil
}
