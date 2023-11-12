package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mikeunge/Tasker/pkg/config"
	"github.com/mikeunge/Tasker/pkg/helpers"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/mikeunge/Tasker/pkg/logger"
)

var db *sql.DB

func Connection() *sql.DB {
	return db
}

func Disconnect() {
	db.Close()
}

func init() {
	var err error

	firstRun := true
	if helpers.FileExists(config.DATABASE_FILE) {
		firstRun = false
	} else {
		log.Debug("Database not found, creating it")
		err = helpers.CreateFile(config.DATABASE_FILE)
		if err != nil {
			log.Error("Could not create database (path: %s), error: %+v", config.DATABASE_FILE, err)
		}
	}

	db, err = sql.Open("sqlite3", config.DATABASE_FILE)
	if err != nil {
		fmt.Printf("Could not open database: %s\n", config.DATABASE_FILE)
		os.Exit(1)
	}
	log.Debug("Successfully connected to db: %s", config.DATABASE_FILE)

	if firstRun {
		err := runMigrations(config.MIGRATIONS_DIR, config.USER_DIR)
		if err != nil {
			log.Error("Could not run migrations, error: %+v", err)
		}
	}
}
