package main

import (
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	_ "time/tzdata"
	"xx.com/yyy/smart-schedule/configuration"
	"xx.com/yyy/smart-schedule/internal/smart-schedule/api"
	"xx.com/yyy/smart-schedule/internal/smart-schedule/builtin"
	"xx.com/yyy/smart-schedule/internal/smart-schedule/database"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "",
				Usage: "-env development/production",
			},
		},
		Action: func(c *cli.Context) error {
			env := c.String("env")
			if env == configuration.EnvProduction {
				configuration.Load(env)
			} else {
				configuration.Load(configuration.EnvDevelopment)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if viper.GetBool("builtin") {
		b := builtin.New()
		b.Register()
	} else {
		dbDriver := database.NewDatabaseDriver()
		apis := api.CreateAPI(dbDriver)
		apis.Register()
	}
}
