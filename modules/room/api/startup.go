package api

import (
	"github.com/vmdt/gogameserver/modules/room/api/configurations"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewRoomDbContext),
		fx.Provide(infrastructure.NewRoomRepository),
		fx.Invoke(func(dbContext *infrastructure.RoomDbContext, log logger.ILogger) {
			db := dbContext.GetModelDB()

			hasTable := db.Migrator().HasTable(&domain.Room{})
			if !hasTable {
				err := db.AutoMigrate(&domain.Room{})
				if err != nil {
					log.Errorf("Failed to migrate room table: %v", err)
				} else {
					log.Info("Room table migrated successfully")
				}
			} else {
				log.Info("Room table already exists, skipping migration")
			}
		}),
		fx.Invoke(configurations.ConfigRoomMediator),
		fx.Invoke(configurations.ConfigEndpoints),
	)
}
