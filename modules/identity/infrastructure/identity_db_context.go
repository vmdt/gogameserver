package infrastructure

import (
	"context"

	"gorm.io/gorm"
)

type IdentityDbContext struct {
	db      *gorm.DB
	context context.Context
}

func NewIdentityDbContext(db *gorm.DB, ctx context.Context) *IdentityDbContext {
	return &IdentityDbContext{
		db:      db,
		context: ctx,
	}
}

func (i *IdentityDbContext) GetModelDB(model interface{}) *gorm.DB {
	return i.db.Model(model)
}
