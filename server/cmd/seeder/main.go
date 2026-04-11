package main

import (
	"debian-ecommerce/internal/data/seeder"
	"debian-ecommerce/pkg/database"
	"debian-ecommerce/pkg/utils"
	"log"
)

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatalf("failed to read file config: %v", err)
	}

	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatalf("failed to connect to postgres database: %v", err)
	}

	logger, err := utils.InitLogger(config.PathLogging, config.Debug)

	// 1. Static master data (safe, fast)
	// err = db.Transaction(func(tx *gorm.DB) error {
	// 	if err := seeder.SeedProvinces(tx, logger); err != nil {
	// 			return err
	// 	}
	// 	if err := seeder.SeedRegencies(tx, logger); err != nil {
	// 			return err
	// 	}
	// 	if err := seeder.SeedDistricts(tx, logger); err != nil {
	// 			return err
	// 	}
	// 	return nil
	// })

	// if err != nil {
	// 	log.Fatalf("transaction location seeder failed: %v", err)
	// }

	// 2. External API sync (NO transaction)
	// if err := seeder.SeedRajaOngkirRegencies(
	// 	db,
	// 	config.RajaOngkirConfig.APIKey,
	// 	logger,
	// ); err != nil {
	// 	log.Fatalf("raja ongkir id seeder failed: %v", err)
	// }
	// if err := seeder.SeedRajaOngkirDistrictsWithLimit(
	// 	db,
	// 	config.RajaOngkirConfig.APIKey,
	// 	logger,
	// 	config.RajaOngkirConfig.APILimit,
	// ); err != nil {
	// 	log.Fatalf("raja ongkir id seeder failed: %v", err)
	// }

	if err := seeder.SeedSubdistricts(db, logger); err != nil {
	// 	return err
	// }

	// if err != nil {
		log.Fatalf("transaction location seeder failed: %v", err)
	}
}
