package main

import "github.com/urfave/cli/v2"

func daemonInstall(ctx *cli.Context) error {
	// TODO: init service, user, config
	return nil
}

func daemonUninstall(ctx *cli.Context) error {
	// TODO: remove all installation stuff
	return nil
}

func daemonRun(ctx *cli.Context) error {
	// TODO: check to see if daemon user, load opts, run
	return nil
}

func daemonStart(ctx *cli.Context) error {
	// TODO: try to use services package to start the daemon
	return nil
}

func daemonStop(ctx *cli.Context) error {
	// TODO: try to use services package to stop the daemon
	return nil
}

func daemonRestart(ctx *cli.Context) error {
	// TODO: try to use services package to restart the daemon
	return nil
}
