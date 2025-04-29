package main

import (
	"github.com/Alexx1088/authservice/internal/db"
	"github.com/Alexx1088/authservice/internal/migrate"
	"log"
)

func main() {
	migrationsPath := "file://migrations"

	if err := db.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Pool.Close()

	dbURL := "postgres://:password@localhost:5433/authservice?sslmode=disable"

	if err := migrate.RunMigrations(migrationsPath, dbURL); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Connected to DB and migration applied successfully.")
}
