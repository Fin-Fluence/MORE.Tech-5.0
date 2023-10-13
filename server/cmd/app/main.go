package main

import (
	"github.com/MORE.Tech-5.0/server/internal/app"
	"github.com/MORE.Tech-5.0/server/pkg/log"
)

func main() {
	logger := log.NewConsoleLogger()

	cfg, err := app.NewConfig()
	if err != nil {
		logger.Fatal("%s", err)
	}

	logger.Fatal("%s", app.Run(cfg, logger))
}
