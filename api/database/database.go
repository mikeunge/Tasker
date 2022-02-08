package database

import (
	"log"
	"os"
	"time"

	"github.com/mikeunge/Tasker/api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	// FIXME: move the logger into it's own constructor, this way we can manage differnt logger for different environments
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	dbString, err := utils.GetEnv("DB_DSN")
	if err != nil {
		return err
	}
	dsn := dbString
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	// Migrate the database
	err = Migrate(DB)
	if err != nil {
		return err
	}

	// Get the underlying sql.DB from gorm so we can change the connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60000) // 1min

	return nil
}
