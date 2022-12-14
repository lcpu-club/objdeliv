// Command objdeliv is the objdeliv server
package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/lcpu-club/objdeliv/server"
	"github.com/lcpu-club/objdeliv/storage"
	_ "github.com/lcpu-club/objdeliv/storage/drivers"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

type configure struct {
	Listen        string                  `yaml:"listen"`
	Driver        string                  `yaml:"driver"`
	DriverOptions storage.DriverConfigure `yaml:"driver-options"`
	DefaultExpire int                     `yaml:"default-expire"`
}

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
			&cli.StringFlag{
				Name:    "default-expire",
				Aliases: []string{"expire", "e"},
				Usage:   "Default object expire time, -1 for unlimited.",
				EnvVars: []string{"OBJDELIV_DEFAULT_EXPIRE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			log.Println()
			log.Println("objdeliv starting...")
			log.Println()
			confText, err := os.ReadFile(ctx.String("config"))
			var listen, driver string = "", ""
			var driverOptions storage.DriverConfigure = nil
			var defaultExpire int
			if err == nil {
				log.Println("Using configure file", ctx.String("config"))
				log.Println()
				conf := &configure{
					DriverOptions: make(storage.DriverConfigure),
				}
				err := yaml.Unmarshal(confText, conf)
				if err != nil {
					return err
				}
				listen = conf.Listen
				driver = conf.Driver
				driverOptions = conf.DriverOptions
				defaultExpire = conf.DefaultExpire
			}
			if err != nil && !os.IsNotExist(err) {
				return err
			}
			if ctx.String("listen") != "" {
				listen = ctx.String("listen")
			}
			if ctx.String("driver") != "" {
				driver = ctx.String("driver")
			}
			if ctx.String("driver-options") != "" {
				json.Unmarshal([]byte(ctx.String("driver-options")), &driverOptions)
			}
			if ctx.Int("default-expire") != 0 {
				defaultExpire = ctx.Int("default-expire")
			}
			log.Println("Storage driver type:	", driver)
			optionText, err := json.Marshal(driverOptions)
			if err != nil {
				return err
			}
			log.Println("Storage driver options:	", string(optionText))
			log.Println("Default expire time:	", strconv.Itoa(defaultExpire)+"s")
			log.Println("Listen address:		", listen)
			log.Println()
			drv, err := storage.NewDriver(driver, driverOptions)
			if err != nil {
				return err
			}
			srv := server.New(listen, drv, defaultExpire)
			return srv.Serve()
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
