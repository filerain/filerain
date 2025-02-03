package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func runMigrations() {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:postgres@localhost:5432/filerain?sslmode=disable")
	if err != nil {
		log.Fatalf("Migrations failed - %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migrations failed - %v", err)
	}
}
