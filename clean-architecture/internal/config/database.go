package config

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config *Config) *gorm.DB {
	newLogger := &GormLogger{
		LogLevel: logger.Info,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable search_path=clean_architecture",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		slog.Error("Failed to connect to database", "details", err.Error())
		os.Exit(1)
	}

	return db
}
