package configurations

import (
	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/modules/identity/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMigrations(dbContext *infrastructure.IdentityDbContext, log logger.ILogger) error {
	db := dbContext.GetModelDB(&domain.User{})
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Error("Failed to run migrations for User model: %v", err)
		return err
	}
	log.Infof("Migrations for User model completed successfully")
	return nil
}
