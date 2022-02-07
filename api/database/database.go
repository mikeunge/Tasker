package database

import (
	"github.com/mikeunge/Tasker/api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbString, err := utils.GetEnv("DB_DSN")
	if err != nil {
		var db gorm.DB
		return &db, err
	}
	dsn := dbString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	// Migrate the database
	err = Migrate(db)
	if err != nil {
		return db, err
	}

	// Get the underlying sql.DB from gorm so we can change the connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60000) // 1min

	return db, nil
}
