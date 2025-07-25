package configurations

import (
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMigrations(dbContext *infrastructure.RoomDbContext, log logger.ILogger) error {
	db := dbContext.GetModelDB(&domain.Room{})
	if err := db.AutoMigrate(&domain.Room{}); err != nil {
		log.Error("Failed to run migrations for Room model: %v", err)
		return err
	}

	db = dbContext.GetModelDB(&domain.RoomPlayer{})
	if err := db.AutoMigrate(&domain.RoomPlayer{}); err != nil {
		log.Error("Failed to run migrations for RoomPlayer model: %v", err)
		return err
	}

	return nil
}
