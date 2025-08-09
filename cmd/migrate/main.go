package main

import (
	"log"
	"os"

	"github.com/Alexx1088/authservice/internal/migrate"
)

func main() {

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	migrationsPath := "file:///app/migrations"

	if err := migrate.RunMigrations(migrationsPath, dbUrl); err != nil {
		log.Fatalf("Migration failed: %v", err)

	}

	log.Println("Migrations complete")
}
