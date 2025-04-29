package main

import (
	"fmt"
	"github.com/Alexx1088/authservice/internal/db"
	"github.com/Alexx1088/authservice/internal/migrate"
	"log"
	"os"
)

func main() {
	migrationsPath := "file://migrations"

	if err := db.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Pool.Close()

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	if err := migrate.RunMigrations(migrationsPath, dbURL); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Connected to DB and migration applied successfully.")
}
