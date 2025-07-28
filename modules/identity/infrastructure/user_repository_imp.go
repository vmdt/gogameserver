package infrastructure

import (
	"context"

	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
)

type UserRepositoryImp struct {
	db      *IdentityDbContext
	generic *postgresgorm.GenericRepository[domain.User]
	ctx     context.Context
}

func NewUserRepositoryImp(db *IdentityDbContext) domain.IUserRepository {
	repo := postgresgorm.NewGenericRepository[domain.User](db.GetModelDB(&domain.User{}))
	return &UserRepositoryImp{
		db:      db,
		generic: repo,
		ctx:     db.context,
	}
}
