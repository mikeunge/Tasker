package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mikeunge/Tasker/pkg/helpers"
)

type IConfig struct {
	DatabasePath string `json:"database_path"`
	LogPath      string `json:"log_path"`
}

const (
	APP_NAME        = "Tasker"
	APP_DESCRIPTION = "A simple task manager."
	APP_VERSION     = "0.0.3"
	APP_AUTHOR      = "@mikeunge"
	APP_DIR         = "/usr/share/tasker"
)

// This are dev variables - they get modified in the init() function
var (
	APP_ENV         = os.Getenv("APP_ENV")
	PWD, _          = os.Getwd()
	USER_CONFIG_DIR = helpers.ExpandPath("~/.config/tasker")
	USER_CONFIG     = USER_CONFIG_DIR + "/config.json"
	USER_DIR        = helpers.ExpandPath("~/.tasker")
	database_path   = USER_DIR + "/database.sqlite"
	log_path        = USER_DIR + "/tasker.log"
	MIGRATIONS_DIR  = PWD + "/migrations"
)

var Config = IConfig{
	DatabasePath: database_path,
	LogPath:      log_path,
}

func init() {
	if helpers.GetStringLength(PWD) <= 0 {
		panic("Could not get current working directory.")
	}

	// production specific config
	if APP_ENV != "developer" {
		MIGRATIONS_DIR = APP_DIR + "/migrations"
	}

	// user specific config
	if !helpers.PathExists(USER_CONFIG_DIR) {
		err := helpers.CreateDirectory(USER_CONFIG_DIR)
		if err != nil {
			return
		}

		// try to create the config file with some default data
		data, _ := json.MarshalIndent(Config, "", " ")
		err = helpers.WriteFile(USER_CONFIG, string(data[:]))
		if err != nil {
			fmt.Printf("Could not create %s\nError: %+v\n", USER_CONFIG, err)
		}
	} else {
		data, err := helpers.ReadFileBytes(USER_CONFIG)
		if err != nil {
			return
		}

		err = json.Unmarshal(data, &Config)
		if err != nil {
			fmt.Printf("Could not parse config, please make sure the config is valid json.\nError: %+v\n", err)
		}
	}

}
