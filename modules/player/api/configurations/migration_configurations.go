package configurations

import (
	"github.com/vmdt/gogameserver/modules/player/domain"
	"github.com/vmdt/gogameserver/modules/player/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMigrations(dbContext *infrastructure.PlayerDbContext, log logger.ILogger) error {
	db := dbContext.GetModelDB()

	if err := db.AutoMigrate(&domain.Player{}); err != nil {
		log.Error("Failed to run migrations for Player model: %v", err)
		return err
	}

	return nil
}
