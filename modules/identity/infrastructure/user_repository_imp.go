package infrastructure

import (
	"context"

	"github.com/pkg/errors"
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
func (repo *UserRepositoryImp) CreateUser(user *domain.User) (*domain.User, error) {
	if err := repo.generic.Add(user, repo.ctx); err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return user, nil
}

func (repo *UserRepositoryImp) GetUserById(id string) (*domain.User, error) {
	var user domain.User
	if err := repo.db.GetModelDB(&user).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (repo *UserRepositoryImp) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := repo.db.GetModelDB(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}
