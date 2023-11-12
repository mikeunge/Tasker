package config

import (
	"os"

	"github.com/mikeunge/Tasker/pkg/helpers"
)

const (
	APP_NAME        = "Tasker"
	APP_DESCRIPTION = "A simple task manager."
	APP_VERSION     = "0.0.3"
	APP_AUTHOR      = "@mikeunge"
	APP_DIR         = "/usr/share/tasker"
)

// This are dev variables - they get modified in the init() function
var (
	APP_ENV        = os.Getenv("APP_ENV")
	PWD, _         = os.Getwd()
	USER_DIR       = helpers.ExpandPath("~/.tasker")
	DATABASE_FILE  = USER_DIR + "/database.sqlite"
	LOG_FILE       = USER_DIR + "/tasker.log"
	MIGRATIONS_DIR = PWD + "/migrations"
)

func init() {
	if helpers.GetStringLength(PWD) <= 0 {
		panic("Could not get current working directory.")
	}

	// production specific config
	if APP_ENV != "developer" {
		MIGRATIONS_DIR = APP_DIR + "/migrations"
	}
}
