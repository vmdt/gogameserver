package infrastructure

import (
	"context"

	"gorm.io/gorm"
)

type BoardGameDbContext struct {
	db      *gorm.DB
	context context.Context
}

func NewBoardGameDbContext(db *gorm.DB, ctx context.Context) *BoardGameDbContext {
	return &BoardGameDbContext{
		db:      db,
		context: ctx,
	}
}

func (b *BoardGameDbContext) GetModelDB(model interface{}) *gorm.DB {
	return b.db.Model(model)
}
