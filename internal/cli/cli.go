package cli

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/mikeunge/Tasker/pkg/config"
)

type Args struct {
	Debug   bool
	Verbose bool
}

func New() (Args, error) {
	parser := argparse.NewParser(config.APP_NAME, config.APP_DESCRIPTION)

	version := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Prints the version"})
	debug := parser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Enable debug logging"})
	verbose := parser.Flag("r", "verbose", &argparse.Options{Required: false, Help: "Logs more (verbose) information"})

	err := parser.Parse(os.Args)
	if err != nil {
		return Args{}, fmt.Errorf("%+v", parser.Usage(err))
	}

	if *version {
		fmt.Printf("v%s\n", config.APP_VERSION)
		os.Exit(0)
	}

	var args = Args{Debug: false, Verbose: false}
	if *debug {
		args.Debug = true
	}

	if *verbose && !*debug {
		args.Verbose = true
	}
	return args, nil
}
