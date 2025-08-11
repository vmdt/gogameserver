package infrastructure

import (
	"context"

	"gorm.io/gorm"
)

type ChatDbContext struct {
	db      *gorm.DB
	context context.Context
}

func NewChatDbContext(db *gorm.DB, ctx context.Context) *ChatDbContext {
	return &ChatDbContext{
		db:      db,
		context: ctx,
	}
}

func (c *ChatDbContext) GetModelDB(model interface{}) *gorm.DB {
	return c.db.Model(model)
}
