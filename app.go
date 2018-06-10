package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/koketani/elapsed/config"
)

type App struct {
	logger *log.Logger
	config *config.Config
}

func NewApp(l *log.Logger, configFile string) (*App, error) {
	if config, err := config.LoadConfig(configFile); err != nil {
		return nil, err
	} else {
		return &App{
			logger: l,
			config: config,
		}, nil
	}
}

func LoadConfig(fileName string) (*Config, error) {
	return &Config{}, nil
}

type Config struct {
}

func (app *App) log() {
	app.logger.Info("test")
}
