package configurations

import (
	"github.com/vmdt/gogameserver/modules/player/domain"
	"github.com/vmdt/gogameserver/modules/player/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMigrations(dbContext *infrastructure.PlayerDbContext, log logger.ILogger) error {
	db := dbContext.GetModelDB()
	hasTable := db.Migrator().HasTable(&domain.Player{})

	if !hasTable {
		if err := db.Migrator().CreateTable(&domain.Player{}); err != nil {
			log.Errorf("Failed to create table: %v", err)
			return err
		}
		log.Infof("Table created successfully")
	} else {
		log.Infof("Table already exists")
	}

	return nil
}
