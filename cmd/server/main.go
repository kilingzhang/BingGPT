package main

import (
	"BingGPT/internal/command"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Name:  "BingGPT",
		Usage: "BingGPT CLI",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Value:   "conf/config.yml",
				Usage:   "load YAML configuration from `FILE`",
				EnvVars: []string{"CONFIG"},
			},
		},
		Commands: []*cli.Command{
			command.Version,
			command.Server,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
