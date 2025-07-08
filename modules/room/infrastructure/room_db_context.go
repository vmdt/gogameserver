package infrastructure

import (
	"context"

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

func (r *RoomDbContext) GetModelDB(model interface{}) *gorm.DB {
	return r.db.Model(model)
}
