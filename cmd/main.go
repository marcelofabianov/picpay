package main

import (
	"log"

	"github.com/marcelofabianov/picpay/config"
	"github.com/marcelofabianov/picpay/internal/infra"
	"github.com/marcelofabianov/picpay/pkg/zap"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	logger, err := zap.NewLogger(cfg.Log)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
	defer logger.Close()

	api := infra.Api(cfg, logger)

	if err := api.Listen(cfg.Api.Address); err != nil {
		logger.Error("error starting server", zap.Error(err))
	}
}
