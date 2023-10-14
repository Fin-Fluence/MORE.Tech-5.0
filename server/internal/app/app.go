package app

import (
	"context"
	"fmt"

	"github.com/MORE.Tech-5.0/server/internal/controller/http"
	"github.com/MORE.Tech-5.0/server/pkg/log"
	"github.com/MORE.Tech-5.0/server/pkg/postgres"
)

func Run(cfg *Config, logger log.Logger) error {
	postgres, err := postgres.New(context.Background(), cfg.Postgres.URI)
	if err != nil {
		return err
	}
	defer postgres.Close()
	logger.Info("database connection established")

	server := http.NewServer(logger, http.ServerOptions{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Origins: cfg.HTTP.Origins,
	})

	logger.Info("server created with address " + server.Addr)
	return fmt.Errorf("server down: %w", server.ListenAndServe())
}
