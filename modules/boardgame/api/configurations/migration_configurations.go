package configurations

import (
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/modules/boardgame/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMigrations(dbContext *infrastructure.BoardGameDbContext, log logger.ILogger) error {
	db := dbContext.GetModelDB(&domain.BattleShip{})
	if err := db.AutoMigrate(&domain.BattleShip{}); err != nil {
		log.Error("Failed to run migrations for BattleShip model: %v", err)
		return err
	}
	log.Infof("Migrations for BattleShip model completed successfully")
	return nil
}
