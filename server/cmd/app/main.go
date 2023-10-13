package main

import (
	"flag"

	"github.com/MORE.Tech-5.0/server/internal/app"
	"github.com/MORE.Tech-5.0/server/pkg/log"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "config file path")
	flag.Parse()

	logger := log.NewConsoleLogger()

	cfg, err := app.NewConfig(cfgPath)
	if err != nil {
		logger.Fatal("%s", err)
	}

	logger.Info("config loaded from: %s", cfgPath)
	logger.Fatal("%s", app.Run(cfg, logger))
}
