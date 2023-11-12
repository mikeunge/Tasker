package cmd

import (
	"fmt"
	"os"

	"github.com/mikeunge/Tasker/internal/cli"
	"github.com/mikeunge/Tasker/internal/services"

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
	projectName := "Project 1"
	projects, err := services.NewProject(projectName)
	if err != nil {
		log.Error("Could not load projects: %+v", err)
	}
	fmt.Printf("%+v\n", projects)

	// TODO: write the UI & the decision tree

	return 0
}
