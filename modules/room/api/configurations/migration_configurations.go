package configurations

import (
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMigrations(dbContext *infrastructure.RoomDbContext, log logger.ILogger) error {
	db := dbContext.GetModelDB(&domain.Room{})
	hasTable := db.Migrator().HasTable(&domain.Room{})

	if !hasTable {
		if err := db.Migrator().CreateTable(&domain.Room{}); err != nil {
			log.Errorf("Failed to create table: %v", err)
			return err
		}
		log.Infof("Table created successfully")
	} else {
		log.Infof("Table already exists")
	}

	hasTable = db.Migrator().HasTable(&domain.RoomPlayer{})

	if !hasTable {
		if err := db.Migrator().CreateTable(&domain.RoomPlayer{}); err != nil {
			log.Errorf("Failed to create table: %v", err)
			return err
		}
		log.Infof("Table created successfully")
	} else {
		log.Infof("Table already exists")
	}

	return nil
}
