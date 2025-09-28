package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// ConnectDB membuat koneksi database yang fleksibel untuk berbagai lingkungan.
func ConnectDB() *gorm.DB {
    // Muat .env secara lokal, abaikan error jika tidak ada (untuk deployment)
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading .env file, using system environment variables.")
    }

    var dsn string

    // Prioritas 1: Coba ambil DSN lengkap dari Railway atau layanan lain
    dsn = os.Getenv("DATABASE_URL")

    // Prioritas 2: Jika DSN tidak ada, coba buat dari variabel Railway standar
    if dsn == "" {
        host := os.Getenv("PGHOST")
        user := os.Getenv("PGUSER")
        pass := os.Getenv("PGPASSWORD")
        name := os.Getenv("PGDATABASE")
        port := os.Getenv("PGPORT")

        if host != "" && user != "" && pass != "" && name != "" && port != "" {
            dsn = fmt.Sprintf(
                "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
                host, user, pass, name, port,
            )
        }
    }

    // Prioritas 3: Jika masih kosong, coba gunakan variabel lokal (fallback)
    if dsn == "" {
        host := os.Getenv("DB_HOST")
        user := os.Getenv("DB_USER")
        pass := os.Getenv("DB_PASS")
        name := os.Getenv("DB_NAME")
        port := os.Getenv("DB_PORT")

        if host != "" && user != "" && pass != "" && name != "" && port != "" {
            dsn = fmt.Sprintf(
                "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
                host, user, pass, name, port,
            )
        }
    }

    // Jika semua cara gagal, hentikan program
    if dsn == "" {
        log.Fatal("Could not find any database connection environment variables.")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("✅ Connected to DB successfully!")
    return db
}