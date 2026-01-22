// Package main runs database migrations for the auth service.
package main

import (
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	databaseURL := os.Getenv("POSTGRES_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	migrationsPath := "file://migrations"

	m, err := migrate.New(
		migrationsPath,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("failed to init migrate: %v", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply")
			return
		}
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("migrations applied successfully")
}
