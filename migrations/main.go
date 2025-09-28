// file: migrate/main.go
package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5" // PENTING: import driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    // Dapatkan URL database dari variabel lingkungan
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        // Fallback ke variabel Railway standar jika tidak ada DATABASE_URL
        dbURL = os.Getenv("PGDATABASE_URL")
    }

    if dbURL == "" {
        log.Fatal("Database URL not found in environment variables.")
    }

    // Jalankan migrasi
    m, err := migrate.New("file://migrations", dbURL)
    if err != nil {
        log.Fatal(err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }

    log.Println("Database migration completed successfully.")
}