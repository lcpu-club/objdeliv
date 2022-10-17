// Command objdeliv is the objdeliv server
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "objdeliv",
		Usage: "Universal temporary large object delivery service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c", "conf"},
				Value:   "./configure.yml",
				Usage:   "Specify the config file path",
				EnvVars: []string{"OBJDELIV_CONFIG"},
			},
			&cli.StringFlag{
				Name:    "listen",
				Aliases: []string{"l", "addr"},
				Usage:   "Listen address",
				EnvVars: []string{"OBJDELIV_LISTEN"},
			},
			&cli.StringFlag{
				Name:    "driver",
				Aliases: []string{"d"},
				Usage:   "Specify the storage driver",
				EnvVars: []string{"OBJDELIV_DRIVER"},
			},
			&cli.StringFlag{
				Name:    "driver-options",
				Aliases: []string{"o"},
				Usage:   "Storage driver options, in JSON",
				EnvVars: []string{"OBJDELIV_DRIVER_OPTIONS"},
			},
		},
		Action: func(ctx *cli.Context) error {
			// TODO: implement the server launching code
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
