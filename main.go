package main

import (
	"fmt"
	"os"

	cmd "github.com/mikeunge/Tasker/cmd/tasker"
	"github.com/mikeunge/Tasker/pkg/config"
	"github.com/mikeunge/Tasker/pkg/helpers"
)

func init() {
	// on every run we check if the USER_DIR exists - if not we try to create it
	if !helpers.PathExists(config.USER_DIR) {
		fmt.Printf("App path does not exist, creating it...\n")
		err := helpers.CreateDirectory(config.USER_DIR)
		if err != nil {
			fmt.Printf("Could not create directory: %s\nError: %+v\n", config.USER_DIR, err)
			os.Exit(-1)
		}
		fmt.Printf("App path: %s\n", config.USER_DIR)
	}
}

func main() {
	os.Exit(cmd.App())
}
