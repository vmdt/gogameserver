package infrastructure

import (
	"context"

	"github.com/vmdt/gogameserver/modules/room/domain"
	"gorm.io/gorm"
)

type RoomDbContext struct {
	db      *gorm.DB
	context context.Context
}

func NewRoomDbContext(db *gorm.DB, ctx context.Context) *RoomDbContext {
	return &RoomDbContext{
		db:      db,
		context: ctx,
	}
}

func (r *RoomDbContext) GetModelDB() *gorm.DB {
	return r.db.Model(&domain.Room{})
}
