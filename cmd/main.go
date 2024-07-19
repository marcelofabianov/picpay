package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/picpay/config"
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

	if err := db.Ping(context.Background()); err != nil {
		logger.Fatal("error pinging database")
	}

	logger.Info("OK")
}
