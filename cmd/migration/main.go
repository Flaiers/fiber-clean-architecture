package main

import (
	"log"

	"github.com/flaiers/fiber-clean-architecture/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(config.MigrateDatabase(cfg.Database.Dsn))
}
