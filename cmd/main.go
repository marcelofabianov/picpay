package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/picpay/config"
	"github.com/marcelofabianov/picpay/internal/infra"
	"github.com/marcelofabianov/picpay/pkg/postgres"
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

	db, err := postgres.Connect(context.Background(), cfg.Db)
	if err != nil {
		logger.Fatal("error connecting to database")
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			logger.Fatal("error closing database connection")
		}
	}()

	api := infra.Api(cfg, logger, db.Conn())

	if err := api.Listen(cfg.Api.Address); err != nil {
		logger.Error("error starting server", zap.Error(err))
	}
}
