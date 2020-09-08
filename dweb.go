package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "dweb",
		Description:          "Centralized service for web hosting with Docker",
		EnableBashCompletion: true,
		Version:              "0.0.1",
		Before:               initLogging,
	}
	app.Authors = append(app.Authors, &cli.Author{
		Name:  "Colton D.",
		Email: "colton@donn.io",
	})
	app.Commands = append(app.Commands, &cli.Command{
		Name: "daemon",
		Subcommands: []*cli.Command{
			{
				Name:   "install",
				Action: daemonInstall,
			},
			{
				Name:   "uninstall",
				Action: daemonUninstall,
			},
			{
				Name:   "run",
				Action: daemonRun,
			},
			{
				Name:   "start",
				Action: daemonStart,
			},
			{
				Name:   "stop",
				Action: daemonStop,
			},
			{
				Name:   "restart",
				Action: daemonRestart,
			},
		},
	})

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func initLogging(ctx *cli.Context) error {
	return nil
}
