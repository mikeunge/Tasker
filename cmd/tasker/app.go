package cmd

import (
	"fmt"
	"os"

	"github.com/mikeunge/Tasker/internal/cli"

	log "github.com/mikeunge/Tasker/pkg/logger"
)

var cliArgs cli.Args

func init() {
	var err error
	cliArgs, err = cli.New()
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(0)
	}

	if cliArgs.Debug {
		log.SetLogLevel("debug")
		log.WriteLogToFile(false)
		log.Info("Tasker running in debug mode")
	} else if cliArgs.Verbose {
		log.SetLogLevel("info")
	}
}

func App() int {
	// TODO: write the UI & the decision tree

	return 0
}
