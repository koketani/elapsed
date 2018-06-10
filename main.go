package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.yml",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		logger := log.New()
		myapp, err := NewApp(logger, c.String("config"))
		if err != nil {
			logger.Errorf("Failed to start: %s", err)
			os.Exit(3)
		}
		myapp.log()
		return nil
	}

	app.Run(os.Args)
}
