package main

import (
	"debian-ecommerce/cmd"
	"debian-ecommerce/internal/data/migration"
	"debian-ecommerce/internal/data/seeder"
	"debian-ecommerce/internal/wire"
	"debian-ecommerce/pkg/database"
	"debian-ecommerce/pkg/utils"
	"log"
)

// @title Debian API
// @version 1.0
// @description E-commerce backend API
// @host localhost:8080
// @BasePath /api/v1
func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatalf("failed to read file config: %v", err)
	}

	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatalf("failed to connect to postgres database: %v", err)
	}

	rdb, err := database.InitRedis(config.Redis)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	logger, err := utils.InitLogger(config.PathLogging, config.Debug)

	// migration
	err = migration.AutoMigrate(db)
	if err != nil {
		log.Println(err)
	}

	// seeder
	err = seeder.SeedAll(db)
	if err != nil {
		log.Println(err)
	}

	route := wire.Wiring(db, rdb, logger, config)
	cmd.APiserver(route)

	// cron scheduler
	route.Scheduler.Start()
}
