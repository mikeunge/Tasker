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
	if helpers.FileExists(config.Config.DatabasePath) {
		firstRun = false
	} else {
		log.Debug("Database not found, creating it")
		err = helpers.CreateFile(config.Config.DatabasePath)
		if err != nil {
			log.Error("Could not create database (path: %s), error: %+v", config.Config.DatabasePath, err)
		}

		// remove "last_migration" if exists, else the new database is empty
		lastMigration := helpers.JoinPath(config.USER_DIR, "last_migration")
		if helpers.FileExists(lastMigration) {
			err = helpers.RemoveFile(lastMigration)
			if err != nil {
				log.Error("Could not remove last_migration: %s, error: %+v", lastMigration, err)
			}
		}
	}

	db, err = sql.Open("sqlite3", config.Config.DatabasePath)
	if err != nil {
		fmt.Printf("Could not open database: %s\n", config.Config.DatabasePath)
		os.Exit(1)
	}
	log.Debug("Successfully connected to db: %s", config.Config.DatabasePath)

	if firstRun {
		err := runMigrations(config.MIGRATIONS_DIR, config.USER_DIR)
		if err != nil {
			log.Error("Could not run migrations, error: %+v", err)
		}
	}
}
