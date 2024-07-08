package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/forbot161602/x-srv-account/source/entry/xvinfo"
	_ "github.com/forbot161602/x-srv-account/source/entry/xvpreset"
	"github.com/forbot161602/x-srv-account/source/entry/xvscript"
)

var (
	app *cli.App
)

func init() {
	app = &cli.App{
		Name:      "srv-account",
		Usage:     "Service: Account",
		Version:   "v1",
		HelpName:  "./main.exe",
		ArgsUsage: "[arguments...]",
		Authors: []*cli.Author{
			&cli.Author{Name: "gordon.lai@starryck.com"},
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			return nil
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:      "show-info",
				Usage:     "Present service information",
				HelpName:  "show-info",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return xvinfo.Execute()
				},
			},
			&cli.Command{
				Name:      "run-script",
				Usage:     "Perform a script",
				HelpName:  "run-script",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return xvscript.Execute()
				},
			},
		},
	}
}

func main() {
	app.Run(os.Args)
}
