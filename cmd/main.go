package main

import (
	"fmt"
	"log"

	"github.com/marcelofabianov/picpay/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println(cfg.Env)
}
