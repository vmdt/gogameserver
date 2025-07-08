package infrastructure

import (
	"context"

	"github.com/vmdt/gogameserver/modules/player/domain"
	"gorm.io/gorm"
)

type PlayerDbContext struct {
	db      *gorm.DB
	context context.Context
}

func NewPlayerDbContext(db *gorm.DB, ctx context.Context) *PlayerDbContext {
	return &PlayerDbContext{
		db:      db,
		context: ctx,
	}
}

func (p *PlayerDbContext) GetModelDB() *gorm.DB {
	return p.db.Model(&domain.Player{})
}
