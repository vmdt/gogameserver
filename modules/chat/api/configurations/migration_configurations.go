package configurations

import (
	"github.com/vmdt/gogameserver/modules/chat/domain"
	"github.com/vmdt/gogameserver/modules/chat/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMigrations(dbContext *infrastructure.ChatDbContext, log logger.ILogger) error {
	db := dbContext.GetModelDB(&domain.Chat{})
	if err := db.AutoMigrate(&domain.Chat{}); err != nil {
		log.Error("Failed to run migrations for Chat model: %v", err)
		return err
	}

	db = dbContext.GetModelDB(&domain.ChatMessage{})
	if err := db.AutoMigrate(&domain.ChatMessage{}); err != nil {
		log.Error("Failed to run migrations for ChatMessage model: %v", err)
		return err
	}
	log.Infof("Migrations for ChatMessage model completed successfully")
	return nil
}
